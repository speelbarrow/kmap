Feature: parse string input
  In order to generate a k-map based on user-inputted arguments
  As a program
  I need to parse said arguments

  @wip
  Scenario: valid input
    When I parse the string "0,3,6,11"
    Then the parsing result should be
      | 0 | 3 | 6 | 11 |

  @wip
  Scenario Outline: variable delimiters
    When I parse the string "<input>"
    Then the parsing result should be
      | <a> | <b> | <c> | <d> |

    Scenarios:
      | input      | a | b | c | d |
      | 8, 3, 1, 7 | 8 | 3 | 1 | 7 |
      | 9:4:2:8    | 9 | 4 | 2 | 8 |
      | 0/1/4/7    | 0 | 1 | 4 | 7 |
      | 3 6 2 9    | 3 | 6 | 2 | 9 |

  @wip
  Scenario Outline: invalid input
    When I parse the string "<input>"
    Then an error should have occurred

    Scenarios:
      | input    |
      # Different delimiters
      | 0, 3,4 5 |
      # Non numeric character
      | 4,f,7,10 |
      # Completely invalid
      | foobar   |
