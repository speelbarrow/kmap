Feature: format output
  In order to provide a useful k-map to the user
  As a program
  I need to format the output nicely

  @wip
  Scenario: size = 2
    Given the k-map size is 2
    And I randomly generate the arguments to the k-map
    When I generate the k-map
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

  @wip
  Scenario: size = 3
    Given the k-map size is 3
    And I randomly generate the arguments to the k-map
    When I generate the k-map
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

  @wip
  Scenario: size = 4
    Given the k-map size is 4
    And I randomly generate the arguments to the k-map
    When I generate the k-map
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