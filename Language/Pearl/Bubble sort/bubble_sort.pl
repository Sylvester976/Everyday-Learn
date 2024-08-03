# Declare an array to be sorted
my @arr = (3, 6, 1, 8, 2);

# Get the length of the array
my $n = scalar @arr;

# Perform the Bubble Sort algorithm
for (my $i = 0; $i < $n - 1; $i++) {
    for (my $j = 0; $j < $n - $i - 1; $j++) {
        # Swap elements if they are in the wrong order
        if ($arr[$j] > $arr[$j+1]) {
            my $temp = $arr[$j];
            $arr[$j] = $arr[$j+1];
            $arr[$j+1] = $temp;
        }
    }
}

# Print the sorted array
print "Sorted array: @arr\n";
