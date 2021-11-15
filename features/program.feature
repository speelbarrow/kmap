Feature: executable program
  In order to generate k-maps for users
  As a developer
  I need to write a program

  Scenario Outline: run the program
    When I run the program
    Then I should be asked "What is the size of the k-map? (3):"
    When I answer "<sizeInput>"
    Then I should be asked "What are the arguments to the k-map?:"
    When I answer ""
    Then I should be asked "What are the don't care conditions of the k-map?:"
    When I answer ""
    Then the program should output an empty k-map of size <size>
    And the program should exit cleanly

    Scenarios:
      | sizeInput | delim | size |
      |           | ","   | 3    |
      | 2         | ", "  | 2    |
      | 3         | " "   | 3    |
      | 4         | ":"   | 4    |

  Scenario Outline: invalid size
    When I run the program
    Then I should be asked "What is the size of the k-map? (3):"
    When I answer "<size>"
    Then I should be asked "Invalid input. Please try again."
    Then I should be asked "What is the size of the k-map? (3):"
    When I answer ""
    Then I should be asked "What are the arguments to the k-map?:"
    When I answer ""
    Then I should be asked "What are the don't care conditions of the k-map?:"
    When I answer ""
    Then the program should output an empty k-map of size 3
    And the program should exit cleanly

    Scenarios:
      | size |
      | 5    |
      | 1    |
