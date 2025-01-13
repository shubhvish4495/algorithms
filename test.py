import math

def NthRoot(n: int, m: int) -> int:
    # Write Your Code Here
    l, h = 1, m

    while l <= h:
        mid = (l + h) // 2
        p_val = math.pow(mid, n)  # Calculate mid^n as a float
        
        if p_val == m:  # Check if mid^n equals m
            return mid
        
        if p_val > m:  # If mid^n is greater than m, reduce the range
            h = mid - 1
        else:  # If mid^n is less than m, increase the range
            l = mid + 1

    return -1