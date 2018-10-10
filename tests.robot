*** Settings ***
Library    SeleniumLibrary

*** Variables ***
${BROWSER}    chrome

*** Keywords ***
End of Test
    Sleep    2
    Close Browser

Open Home page
    Open Browser    https://localhost:4433    ${BROWSER}
    Maximize Browser Window

Open Login Window
    Click Element    login

*** Test Cases ***
Test Open Home Page
    Open Home page
    Open Login Window
    End of Test