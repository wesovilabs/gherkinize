Feature: Step with more characters than expected
  Scenario: eat 5 out of 12
  Given there are 12 cucumbers
  When I eat 5 cucumbers
  Then I should have 7 cucumbers

  Scenario: eat 5 out of 20
  Given there are 20 cucumbers which were bought in the supermarket next to home.
  When I eat 5 cucumbers
  Then I should have 15 cucumbers