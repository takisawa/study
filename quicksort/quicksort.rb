
def quicksort(ary)
  if ary.size <= 1
    return ary
  end

  pivot = ary.shift

  less_than_or_equal_to = []
  greater_than = []

  ary.each do |elem|
    if elem <= pivot
      less_than_or_equal_to << elem
    else
      greater_than << elem
    end
  end

  quicksort(less_than_or_equal_to) + [pivot] + quicksort(greater_than)
end

p quicksort([4, 5, 2, 1, 3])
