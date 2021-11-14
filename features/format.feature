@wip
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
        | 0 | 0 |
        ---------
      x | 0 | 0 |
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
        | 0 | 0 | 0 | 0 |
        -----------------
      x | 0 | 0 | 0 | 0 |
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
        | 0 | 0 | 0 | 0 |
        -----------------
        | 0 | 0 | 0 | 0 |
        ----------------- x
        | 0 | 0 | 0 | 0 |
      w -----------------
        | 0 | 0 | 0 | 0 |
        -----------------
                z
      """

  Scenario: with input
    Given the arguments to the k-map are
      | 0 | 4 | 6 |
    And the don't care conditions of the k-map are
      | 1 | 2 | 5 |
    When I initialize the k-map
    And I create the output for the generated k-map
    Then the formatted output should match
      """
                    y
        -----------------
        | 1 | X | 0 | X |
        -----------------
      x | 1 | X | 0 | 1 |
        -----------------
                z
      """