<?hh

namespace RocketLabs\SimpleLinearRegression;

type Coefficients = shape('w0' => float, 'w1' => float);

final class Model
{
  public function __construct(private Coefficients $coefficients)
  {
  }

  public function f(float $input): float
  {
    return $this->coefficients['w0'] + ($this->coefficients['w1'] * $input);
  }
}
