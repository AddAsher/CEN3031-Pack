describe('My First Test', () => {
  it('Gets, types and asserts', () => {
    cy.visit('http://localhost:4200')

    cy.contains('PACK')

    // Get an input, type into it
    cy.get('username').type('hello')

    //  Verify that the value has been updated
    // cy.get('.action-email').should('have.value', 'fake@email.com')
  })
})