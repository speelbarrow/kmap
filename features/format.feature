Feature: format output
  In order to provide a useful k-map to the user
  As a program
  I need to format the output nicely

  Scenario: size = 2
    Given the k-map size is 2
    When I initialize the k-map
    And I create the output for the generated k-map
    Then the formatted output should match
      """
              y
        ---------
        | X | X |
        ---------
      x | X | X |
        ---------
      """

  Scenario: size = 3
    Given the k-map size is 3
    When I initialize the k-map
    And I create the output for the generated k-map
    Then the formatted output should match
      """
                    y
        -----------------
        | X | X | X | X |
        -----------------
      x | X | X | X | X |
        -----------------
                z
      """

  Scenario: size = 4
    Given the k-map size is 4
    When I initialize the k-map
    And I create the output for the generated k-map
    Then the formatted output should match
      """
                    y
        -----------------
        | X | X | X | X |
        -----------------
        | X | X | X | X |
        ----------------- x
        | X | X | X | X |
      w -----------------
        | X | X | X | X |
        -----------------
                z
      """