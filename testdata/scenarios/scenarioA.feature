Feature: Scenario with no assertion statement (then)

  Scenario: eat 5 out of 12
  Given there are 12 cucumbers
  When I eat 5 cucumbers
  Then I should have 7 cucumbers


  Scenario: eat 5 out of 20
  Given there are 20 cucumbers
  When I eat 5 cucumbers
  And I should have 15 cucumbers








