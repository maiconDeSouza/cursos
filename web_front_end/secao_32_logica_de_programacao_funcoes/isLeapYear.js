function isLeapYear(year){
    let result = true

    if (year % 4 !== 0) result = false

    if (year % 100 === 0) result = false

    if (year % 400 === 0) result = true


    return result
    
}

console.log(isLeapYear(4))

function isLeapYearV2(year) {
    // Um ano é bissexto se:
    // (Divisível por 4 E não divisível por 100) OU (Divisível por 400)
    return (year % 4 === 0 && year % 100 !== 0) || (year % 400 === 0);
}
