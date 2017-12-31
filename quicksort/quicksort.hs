
quicksort [] = []
quicksort [x] = [x]
quicksort (pivot:xs) = quicksort [x | x <- xs, x <= pivot] ++ [pivot] ++ quicksort [x | x <- xs, x > pivot]

main = putStrLn $ show $ quicksort [4, 5, 2, 1, 3]
