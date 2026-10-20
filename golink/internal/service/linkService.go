package service

import (
	"errors"
	"golink/internal/config"
	errapp "golink/internal/errApp"
	"golink/internal/model"
	"golink/internal/repository"
	"uuid"

	"github.com/go-playground/validator/v10"
)

type LinkService struct {
	repo     repository.LinkRepositoryInterface
	env      *config.ENV
	validate *validator.Validate
}

func NewLinkService(repo repository.LinkRepositoryInterface, env *config.ENV) *LinkService {
	return &LinkService{repo: repo, env: env, validate: validator.New()}
}

func (s *LinkService) CreateLink(userID uuid.UUID, linkRequest model.LinkRequest) (*model.LinkResponse, *errapp.ErrResponse) {
	if err := s.validate.Struct(linkRequest); err != nil {
		return nil, errapp.NewBadRequestError("Problema em um dos campos: "+err.Error(), err)
	}

	totalLinks, err := s.repo.CountByUserID(userID)
	if err != nil {
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}
	if totalLinks >= 5 {
		return nil, errapp.NewBadRequestError("Total de links exceditos, são permitidos ate 5", nil)
	}

	link := linkRequestToLink(linkRequest, userID)

	if err := s.repo.Create(*link); err != nil {
		if errors.Is(err, errapp.ErrTitleAlreadyInLink) {
			return nil, errapp.NewBadRequestError("Problema com titulo do link: "+err.Error(), err)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	linkResponse := linkToLinkResponse(*link)

	return linkResponse, nil
}

func (s *LinkService) Redirect(title string) (*model.LinkResponse, *errapp.ErrResponse) {
	link := model.Link{}
	if err := s.repo.FindByTitle(title, &link); err != nil {
		if errors.Is(err, errapp.ErrLinkNotFound) {
			return nil, errapp.NewNotFoundError("Página não encontrada", err)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}
	linkResponse := linkToLinkResponse(link)

	return linkResponse, nil
}

func (s *LinkService) Update(id uuid.UUID, linkUpdate *model.LinkUpdate) (*model.LinkResponse, *errapp.ErrResponse) {
	link := model.Link{}
	if err := s.repo.FindByID(id, &link); err != nil {
		if errors.Is(err, errapp.ErrLinkNotFound) {
			return nil, errapp.NewNotFoundError("Página não encontrada", err)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	if len(linkUpdate.Title) > 0 {
		if len(linkUpdate.Title) < 3 || len(linkUpdate.Title) > 50 {
			return nil, errapp.NewBadRequestError("O Título deve ter mais do que duas letras e menos que 51", nil)
		}
		link.Title = linkUpdate.Title
	}

	if linkUpdate.URL != "" {
		if err := s.validate.Var(linkUpdate.URL, "required,url"); err != nil {
			return nil, errapp.NewBadRequestError("URL inválida", nil)
		}
		link.URL = linkUpdate.URL
	}

	if err := s.repo.Update(id, &link); err != nil {
		if errors.Is(err, errapp.ErrLinkNotFound) {
			return nil, errapp.NewNotFoundError("Página não encontrada", err)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	linkResponse := linkToLinkResponse(link)

	return linkResponse, nil
}
