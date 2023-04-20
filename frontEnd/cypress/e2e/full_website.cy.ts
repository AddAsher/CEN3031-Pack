import{slowCypressDown} from 'cypress-slow-down'
slowCypressDown()
describe('Full Website Test', () => {
  it('Visits the initial project page', () => {
    cy.visit('/')
    cy.get('#username').type('Al Gator')
    cy.get('#password').type('gatorgator')
    cy.get('#Login').click()
  })
})
