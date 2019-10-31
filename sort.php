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

function sortFile(string $path) {
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

    quickSort($nums);
}

define('ITERS', 2);

function benchTest(string $file) {
    $start = microtime(true);
    for ($i = 0; $i < ITERS; $i++) {
        sortFile($file);
    }
    $total = round(microtime(true) - $start, 5);
    $avg = round($total / ITERS, 5);

    return $file . ': total = ' . $total . 's; average = ' . $avg . 's';
}

echo benchTest('./nums_1000') . PHP_EOL;
echo benchTest('./nums_10000') . PHP_EOL;
echo benchTest('./nums_100000') . PHP_EOL;
echo benchTest('./nums_1000000') . PHP_EOL;
echo benchTest('./nums_10000000') . PHP_EOL;
