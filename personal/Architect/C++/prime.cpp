#include<iostream>
using namespace std;

bool is_prime(int n) {
   if (n < 2) {
      return false;
   }
   for (int i = 2; i*i <= n; i++) {
      if (n % i == 0) {
         return false;
      }
   }
   return true;
}

int main() {
   int limit = 1000;
   for (int i = 3; i <= limit; i++) {
      if (is_prime(i)) {
         cout << i << " is prime!\n";
      }
   }
   return 0;
}

