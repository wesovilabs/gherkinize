aFeature: This is a feature

  Scenario: Coche arrancando
    When arranca
    Given a coche
    Then frena

  Scenario: Coche arrancando 2
    Given a coche
    When arranca
    Then frena
