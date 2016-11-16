<?hh

namespace RocketLabs\SimpleLinearRegression;

require_once 'SimpleLinearRegression.hh';
require_once 'Model.hh';

type Observation = Map<string, float>;
type DataSet = Vector<Observation>;

(function(): void {
  $dataSet = Vector {
    Map{'sqmt' => 123, 'price' => 302030},
    Map{'sqmt' => 86, 'price' => 123000},
    Map{'sqmt' => 45, 'price' => 68000},
    Map{'sqmt' => 200, 'price' => 400876},
    Map{'sqmt' => 66, 'price' => 110423},
    Map{'sqmt' => 90, 'price' => 120432},
  };

  $simpleLinearRegression = new SimpleLinearRegression($dataSet, 'sqmt', 'price');

  while(true) {
    $line = trim(readline("Enter square meters of the house: "));
    $input = (float) $line;
    $output = $simpleLinearRegression->predict($input);

    echo sprintf("Predicted selling price: € %.2f\n", $output);
  }
})();
