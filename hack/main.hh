<?hh

namespace RocketLabs\SimpleLinearRegression;

require_once 'SimpleLinearRegression.hh';
require_once 'Model.hh';
require_once 'DataSet.hh';

(function(): void {
  $dataSet = DataSet::fromCsvFile('../dataset.csv');

  $simpleLinearRegression = new SimpleLinearRegression($dataSet, 'Square meters', 'Price');

  while(true) {
    $line = trim(readline("Enter square meters of the house: "));
    $input = (float) $line;
    $output = $simpleLinearRegression->predict($input);

    echo sprintf("Predicted selling price: € %.2f\n", $output);
  }
})();
