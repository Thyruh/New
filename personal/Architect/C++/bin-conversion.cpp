#include <iostream>
#include <bitset>

int main() {
    int num = 483934;
    std::cout << "The binary of num variable is: " << std::bitset<8>(num) << '\n';
    std::cout << "The decimal of num variable is: " << num << '\n';
    return 0;
}

