#User function Template for python3

class Solution:
    # Function used for sorting jobs according to their deadlines
    def JobSequencing(self, id, deadline, profit):
        #code here
        if not (len(job_ids) == len(deadlines) == len(profits)):
            return 0
    
        # Pair the jobs with their deadlines and profits
        jobs = list(zip(job_ids, deadlines, profits))
    
        # Sort the jobs by deadline; if deadlines are the same, sort by profit descending
        jobs.sort(key=lambda x: (x[1], -x[2]))
    
        # Unpack sorted jobs
        job_ids, deadlines, profits = zip(*jobs)
    
        # Initialize variables
        max_profit = profits[0]
        prev_deadline = deadlines[0]
        curr_profit = profits[0]
    
        # Iterate through the jobs
        for i in range(1, len(deadlines)):
            if deadlines[i] == prev_deadline:
                if curr_profit < profits[i]:
                    max_profit -= curr_profit
                    curr_profit = profits[i]
                    max_profit += curr_profit
            else:
                max_profit += profits[i]
                prev_deadline = deadlines[i]
                curr_profit = profits[i]
    
        return max_profit


#{ 
 # Driver Code Starts
#Initial Template for Python 3
import atexit
import io
import sys
import heapq

if __name__ == "__main__":
    t = int(input())
    for _ in range(t):
        job_ids = list(map(int, input().split()))
        deadlines = list(map(int, input().split()))
        profits = list(map(int, input().split()))
        obj = Solution()
        ans = obj.JobSequencing(job_ids, deadlines, profits)
        print(ans[0], ans[1])
        print("~")

# } Driver Code Ends