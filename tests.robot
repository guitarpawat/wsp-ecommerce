*** Settings ***
Library    SeleniumLibrary

*** Test Cases ***
Test Open
    Open Browser    https://www.google.com    chrome
    Close Browser