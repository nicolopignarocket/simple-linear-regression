<?hh

namespace RocketLabs\SimpleLinearRegression;

type Observation = Map<string, float>;

final class DataSet
{
  private function __construct(private Vector<string> $columns, private Vector<Observation> $rows)
  {
  }

  public static function fromCsvFile($filePath)
  {
    $csvRows = explode("\n", file_get_contents($filePath));
    $columns = new Vector(explode(",", array_shift($csvRows)));

    $rows = Vector{};

    foreach ($csvRows as $row) {
      $cells = explode(",", $row);
      $observation = Map{};
      $validRow = true;

      foreach ($cells as $key => $cell) {
        if ($cell === "") {
          $validRow = false;
          continue;
        }

        $observation[$columns[$key]] = floatval($cell);
      }

      if ($validRow) {
        $rows[] = $observation;
      }
    }

    return new self($columns, $rows);
  }

  public function getRows(): Vector<Observation>
  {
    return $this->rows;
  }

  public function apply((function(mixed): bool) $fn): DataSet {
    $rows = array_filter($this->rows, $fn);

    return new DataSet(clone $this->columns, new vector($rows));
  }
}
