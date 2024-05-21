package main

func asteroidCollision(asteroids []int) (ans []int) {

	if len(asteroids) == 0 {
		return ans
	}
	curr := 0
	for curr < len(asteroids) {
		if len(ans) > 0 && ans[len(ans)-1] > 0 && asteroids[curr] < 0 {

			if ans[len(ans)-1] == -asteroids[curr] {
				ans = ans[:len(ans)-1]
				curr++
			} else if ans[len(ans)-1] > -asteroids[curr] {
				curr++
			} else {
				ans = ans[:len(ans)-1]
			}
		} else {
			ans = append(ans, asteroids[curr])
			curr++
		}
	}
	return ans
}

//class Solution {
//public int[] asteroidCollision(int[] nums) {
//
//if (nums.length == 0)
//return new int[]{};
//
//Stack<Integer> stack = new Stack<>();
////        stack.push(nums[0]);
//
//int curr = 0;
//while (curr < nums.length) {
//
//// case of collission
//if (stack.size() > 0 && stack.peek() > 0 && nums[curr] < 0) {
//// left dies
//if (stack.peek() < Math.abs(nums[curr])) {
//stack.pop();
//} // both dies
//else if (stack.peek() == Math.abs(nums[curr])) {
//stack.pop();
//curr++;
//} else {
//// right dies
//curr++;
//}
//} else {
//// no collission
//stack.push(nums[curr]);
//curr++;
//}
//}
//
//int[] ans = new int[stack.size()];
//Iterator<Integer> iterator = stack.iterator();
//int i = 0;
//while(iterator.hasNext()){
//ans[i++] = iterator.next();
//}
//
//return ans;
//}
//}
