Feature: kmap type
  In order to manipulate a k-map
  As a program
  I need to understand what a k-map is

  Scenario Outline: arguments
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
  Scenario Outline: don't care conditions
    Given the k-map size is <size>
    And I randomly generate the don't care conditions of the k-map
    When I initialize the k-map
    Then the k-map values should match the don't care conditions

    Scenarios:
      | size |
      | 2    |
      | 3    |
      | 4    |

  @wip
  Scenario Outline: arguments and don't care conditions
    Given the k-map size is <size>
    And I randomly generate the arguments and don't care conditions for the k-map
    When I initialize the k-map
    Then the k-map values should match the arguments and don't care conditions

    Scenarios:
      | size |
      | 2    |
      | 3    |
      | 4    |

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

  Scenario Outline: invalid size
    Given the k-map size is <size>
    When I initialize the k-map
    Then an error should have occurred

    Scenarios:
      | size |
      | 5    |
      | 1    |

  @wip
  Scenario: invalid arguments and don't care conditions
    Given the k-map size is 3
    And the arguments to the k-map are
      | 1 | 3 | 7 |
    And the don't care conditions of the k-map are
      | 2 | 3 | 5 |
    When I initialize the k-map
    Then an error should have occurred
