declare function log(n: i32): void;

export function minusOne(n: i32): i32 {
  log(n);
  return n - 1;
}

export function fizzbuzz(n: i32): String | null {
  if (n % 15 == 0) {
    return "FizzBuzz";
  } else if (n % 3 == 0) {
    return "Fizz";
  } else if (n % 5 == 0) {
    return "Buzz";
  }

  return null;
}

memory.grow(2);
store<u8>(0, 21);
store<u8>(1, 99);

export function readMemory(n: i32): i32 {
  return load<u8>(n);
}

