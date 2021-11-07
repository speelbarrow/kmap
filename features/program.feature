Feature: executable program
  In order to generate k-maps for users
  As a developer
  I need to write a program

  @wip
  Scenario Outline: run the program
    When I run the program
    Then I should be asked "What is the size of the k-map? (3): "
    When I answer "<size>"
    Then I should be asked "What are the arguments to the k-map?: "
    When I randomly generate the arguments to the k-map
    And I answer the randomly generated arguments seperated by <delim>
    Then the program should output a proper k-map

    Scenarios:
      | size | delim |
      |      | ","   |
      | 2    | ", "  |
      | 3    | " "   |
      | 4    | ":"   |
