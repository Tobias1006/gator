*** Settings ***
Documentation   This test case will test the login functionality of the 
...             gator CLI application.
Library          Process

*** Variables ***
${COMMAND}       go
${WORKDIR}       ../gatorgo/
${USER1}         schorschi
${USER2}         SchOrSchi
${USER3}         SCHORSCHI
${USER4}         5C#OR5C#!

*** Test Cases ***

Login - no username
    [Documentation]    Starts gatorand checks the error message when too few arguments are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    login    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    1
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    incorrect number of arguments for the command found

Login - Happy path_Username1
    [Documentation]    Starts gator and checks the result when a known command and username are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    login    ${USER1}    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    User   ${USER1}   has been set successfully.

Login - Happy path_Username2
    [Documentation]    Starts gator and checks the result when a known command and username are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    login    ${USER2}    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    User   ${USER2}   has been set successfully.

Login - Happy path_Username3
    [Documentation]    Starts gator and checks the result when a known command and username are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    login    ${USER3}    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    User   ${USER3}   has been set successfully.

Login - Happy path_Username4
    [Documentation]    Starts gator and checks the result when a known command and username are supplied

    ${result}=    Run Process    ${COMMAND}    run    .    login    ${USER4}    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    User   ${USER4}   has been set successfully.
