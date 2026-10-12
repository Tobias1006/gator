*** Settings ***
Documentation   This test case will test the reset functionality of the 
...             gator CLI application.
Library          Process

*** Variables ***
${COMMAND}       go
${WORKDIR}       ../gatorgo/
${USER1}         Lollo
${USER2}         Gisela
${USER3}         schorschi
${USER4}         SCHORSCHI

*** Test Cases ***

Users - Happy Path
    [Documentation]    Starts gator and checks the error message when too few arguments are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    users    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    ${USER1}
    Should Contain    ${full_output}    ${USER2}
    Should Contain    ${full_output}    ${USER3}
    Should Contain    ${full_output}    ${USER4} (current)

Users - Too many arguments
    [Documentation]    Checks if gator reset function is dependent on supplied arguments
    
    ${result}=    Run Process    ${COMMAND}    run    .    users  ${USER1}  ${USER1}   cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    ${USER1}
    Should Contain    ${full_output}    ${USER2}
    Should Contain    ${full_output}    ${USER3}
    Should Contain    ${full_output}    ${USER4} (current)
