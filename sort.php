<?php 

ini_set('memory_limit', '8G');

function quickSort(array $items): array {
    if (count($items) < 2) {
        return $items;
    }
    $control = array_shift($items);
    $left  = [];
    $right = [];
    foreach ($items as $item) {
        if ($item > $control) {
            $right[] = $item;
        }
        else {
            $left[] = $item;
        }
    }
    $items = array_merge(
        quickSort($left),
        [$control],
        quickSort($right)
    );
    return $items;
}

function sortFile(string $path, bool $native = false) {
    $h = fopen($path, 'r');
    $nums = [];
    while (true) {
        $row = fgets($h);
        if (false === $row) {
            break;
        }

        $row = trim($row);

        $nums[] = (int)$row;
    }

    fclose($h);

    if (!$native) {
        quickSort($nums);

        return;
    }

    sort($nums);
}

define('ITERS', 2);

function benchTest(string $file, bool $native) {
    $start = microtime(true);
    for ($i = 0; $i < ITERS; $i++) {
        sortFile($file, $native);
    }
    $total = round(microtime(true) - $start, 5);
    $avg = round($total / ITERS, 5);

    return ($native ? 'native sort' : 'quickSort') . ' for ' . $file . ': total = ' . $total . 's; average = ' . $avg . 's';
}

echo benchTest('./nums_1000', false) . PHP_EOL;
echo benchTest('./nums_10000', false) . PHP_EOL;
echo benchTest('./nums_100000', false) . PHP_EOL;
echo benchTest('./nums_1000000', false) . PHP_EOL;
echo benchTest('./nums_10000000', false) . PHP_EOL;

echo benchTest('./nums_1000', true) . PHP_EOL;
echo benchTest('./nums_10000', true) . PHP_EOL;
echo benchTest('./nums_100000', true) . PHP_EOL;
echo benchTest('./nums_1000000', true) . PHP_EOL;
echo benchTest('./nums_10000000', true) . PHP_EOL;