<?hh

namespace RocketLabs\SimpleLinearRegression;

newtype RssGradient = shape('dw0' => float, 'dw1' => float, 'magnitude' => float);

final class SimpleLinearRegression
{
  private Model $model;

  public function __construct(private DataSet $dataset, private string $inputColumn, private string $outputColumn)
  {
    $this->buildModel();
  }

  public function predict(float $input): float
  {
    return $this->model->f($input);
  }

  private function buildModel(): void
  {
    echo "Finding coefficients...\n";
    $acceptedErrorThreshold = 1.;
    $stepSize = 1.;
    $estimatedCoefficients = $previousCoefficients = shape('w0' => 0., 'w1' => 0.);
    $iterations = 0;
    $previousMagnitude = -1.;

    while (true) {
      $dRss = $this->dRss($estimatedCoefficients);

      if ($dRss['magnitude'] < $acceptedErrorThreshold) {
        break;
      }

      if ($previousMagnitude >= 0) {
        if ($dRss['magnitude'] > $previousMagnitude) {
          $estimatedCoefficients = $previousCoefficients;
          $stepSize *= .5;
        } else {
          $previousMagnitude = $dRss['magnitude'];
          $stepSize *= 1.05;
        }
      } else {
        $previousMagnitude = $dRss['magnitude'];
      }

      echo sprintf("Iteration %d: est. w0 %g, est. w1 %g, dRssW0 %g, dRssW1 %g, magnitude %g\n",
			             $iterations, $estimatedCoefficients['w0'], $estimatedCoefficients['w1'], $dRss['dw0'], $dRss['dw1'], $dRss['magnitude']);

      $previousCoefficients = $estimatedCoefficients;

      $estimatedCoefficients = shape(
        'w0' => $estimatedCoefficients['w0'] - ($stepSize * $dRss['dw0']),
        'w1' => $estimatedCoefficients['w1'] - ($stepSize * $dRss['dw1']),
      );

      $iterations++;

      echo sprintf("Next step size: %g\n", $stepSize);
    }

    echo sprintf("Optimal coefficients found: w0 = %f, w1 = %f\n", $estimatedCoefficients['w0'], $estimatedCoefficients['w1']);

    $this->model = new Model($estimatedCoefficients);
  }

  private function dRss(Coefficients $coefficients): RssGradient
  {
    $sumW0 = $sumW1 = 0.;

    foreach ($this->dataset->getRows() as $observation) {
      $xi = $observation[$this->inputColumn];
      $yi = $observation[$this->outputColumn];

      $partialTerm = $yi - $coefficients['w0'] - ($coefficients['w1'] * $xi);
      $sumW0 += $partialTerm;
      $sumW1 += $partialTerm * $xi;
    }

    $dw0 = -2 * $sumW0;
    $dw1 = -2 * $sumW1;
    $magnitude = sqrt(pow($dw0, 2) + pow($dw1, 2));

    return shape(
      'dw0' => $dw0,
      'dw1' => $dw1,
      'magnitude' => $magnitude
    );
  }
}
