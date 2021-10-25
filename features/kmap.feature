Feature: kmap type
  In order to manipulate a k-map
  As a program
  I need to understand what a k-map is

  @wip
  Scenario Outline: values
    Given the k-map size is <size>
    And I randomly generate the arguments to the k-map
    When I initialize the k-map
    Then the k-map values should match the arguments

    Scenarios:
      | size |
      | 2    |
      | 3    |
      | 4    |

  @wip
  Scenario Outline: size, rows, columns
    Given the k-map size is <size>
    When I initialize the k-map
    Then the "Size" property of the k-map should be <size>
    And the "Rows" property of the k-map should be <row>
    And the "Cols" property of the k-map should be <col>

    Scenarios:
      | size | row | col |
      | 2    | 2   | 2   |
      | 3    | 2   | 4   |
      | 4    | 4   | 4   |

  @wip
  Scenario: in order of minterms
    Given the k-map size is 4
    And the arguments to the k-map are
      | 0 | 3 | 5 | 6 | 9 | 10 | 12 | 15 |
    When I initialize the k-map
    Then the k-map values should match
      | 1 | 0 | 1 | 0 |
      | 0 | 1 | 0 | 1 |
      | 0 | 1 | 0 | 1 |
      | 1 | 0 | 1 | 0 |
    And the Minterms method should match
      | 1 | 0 | 0 | 1 |
      | 0 | 1 | 1 | 0 |
      | 0 | 1 | 1 | 0 |
      | 1 | 0 | 0 | 1 |
