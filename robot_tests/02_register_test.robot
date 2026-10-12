*** Settings ***
Documentation   This test case will test the register functionality of the 
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

Register - no username
    [Documentation]    Starts gatorand checks the error message when too few arguments are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    register    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    1
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    invalid number of arguments provided

Register - Happy path_Username1
    [Documentation]    Starts gator and checks the result when a known command and username are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    register    ${USER1}    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    U-Name: ${USER1}
    Should Contain    ${full_output}    U-CreatedAt
    Should Contain    ${full_output}    U-UpdatedAt
    Should Contain    ${full_output}    U-ID:

Register - Happy path_Username2
    [Documentation]    Starts gator and checks the result when a known command and username are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    register    ${USER2}    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    U-Name: ${USER2}
    Should Contain    ${full_output}    U-CreatedAt
    Should Contain    ${full_output}    U-UpdatedAt
    Should Contain    ${full_output}    U-ID:

Register - Happy path_Username3
    [Documentation]    Starts gator and checks the result when a known command and username are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    register    ${USER3}    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    U-Name: ${USER3}
    Should Contain    ${full_output}    U-CreatedAt
    Should Contain    ${full_output}    U-UpdatedAt
    Should Contain    ${full_output}    U-ID:

Register - Happy path_Username4
    [Documentation]    Starts gator and checks the result when a known command and username are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    register    ${USER4}    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    0
    
    ${full_output}=    Catenate    ${result.stdout}    ${result.stderr}
    Should Contain    ${full_output}    U-Name: ${USER4}
    Should Contain    ${full_output}    U-CreatedAt
    Should Contain    ${full_output}    U-UpdatedAt
    Should Contain    ${full_output}    U-ID:

Register - Duplicate registry
    [Documentation]    Starts gator and checks the result when a known command and username are supplied
    
    ${result}=    Run Process    ${COMMAND}    run    .    register    ${USER1}    cwd=${WORKDIR}
    
    Should Be Equal As Integers    ${result.rc}    1
