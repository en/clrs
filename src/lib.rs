#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {}
}

#[cfg(test)]
mod test_cases {
    pub static SORT: (([i32; 6], [i32; 6]),
     ([char; 8], [char; 8]),
     ([u64; 0], [u64; 0]),
     ([f32; 1], [f32; 1]),
     ([i32; 6], [i32; 6]),
     ([i32; 6], [i32; 6]),
     ([i32; 6], [i32; 6]),
     ([i32; 13], [i32; 13])) =
        (([5, 2, 4, 6, 1, 3], [1, 2, 3, 4, 5, 6]),
         (['y', 'u', 'a', 'n', 'c', 'h', 'a', 'o'], ['a', 'a', 'c', 'h', 'n', 'o', 'u', 'y']),
         ([], []),
         ([1.0], [1.0]),
         ([31, 41, 59, 26, 41, 58], [26, 31, 41, 41, 58, 59]),
         ([1, 2, 3, 4, 5, 6], [1, 2, 3, 4, 5, 6]),
         ([6, 5, 4, 3, 2, 1], [1, 2, 3, 4, 5, 6]),
         ([74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586],
          [-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845]));
    pub static SEARCH: [i32; 5] = [1, 2, 3, 4, 5];
    pub struct SearchV {
        pub v: i32,
        pub r: Result<usize, i32>,
    }
    pub static SEARCH_RESULT: [SearchV; 7] = [SearchV { v: 0, r: Err(-1) },
                                              SearchV { v: 1, r: Ok(0) },
                                              SearchV { v: 2, r: Ok(1) },
                                              SearchV { v: 3, r: Ok(2) },
                                              SearchV { v: 4, r: Ok(3) },
                                              SearchV { v: 5, r: Ok(4) },
                                              SearchV { v: 6, r: Err(-1) }];
}

mod ch01;
mod ch02;
