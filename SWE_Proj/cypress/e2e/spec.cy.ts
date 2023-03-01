import{slowCypressDown} from 'cypress-slow-down'

slowCypressDown()

describe('My First Test', () => {
  it('Gets, types and asserts', () => {
    cy.visit('http://localhost:4200')
    
    cy.contains('PACK')

    // Get an input, type into it
    cy.get('#username').type('hello')
    cy.get('#password').type('world')

    // Click the button to save values
    cy.get('#button').click()
  })
})