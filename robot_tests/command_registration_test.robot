*** Settings ***
Documentation   This test case will test the login functionality of the 
...             gator CLI application.
Library          Process

*** Variables ***
${COMMAND}       go
${WORKDIR}       ../gatorgo/

*** Test Cases ***

Login - no command
    [Documentation]    Starts gator and checks the error message when no argument is supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    1
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    Too few arguments.
    
Login - too many arguments
    [Documentation]    Starts gator and checks the error message when too many arguments are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    login    login    login    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    1
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    Incorrect number of arguments for the command found.

Login - unknown argument
    [Documentation]    Starts gator and checks the error message when an unknown argument is supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    username    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    1
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    Unknown command.