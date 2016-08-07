#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {}
}

#[cfg(test)]
mod test_cases {
    use ch02::merge_sort::Value;

    // Thank God there are vim macros...
    pub static SORT: (([Value<char>; 8], [Value<char>; 8]),
     ([Value<i32>; 6], [Value<i32>; 6]),
     ([Value<u64>; 0], [Value<u64>; 0]),
     ([Value<f32>; 1], [Value<f32>; 1]),
     ([Value<i32>; 6], [Value<i32>; 6]),
     ([Value<i32>; 6], [Value<i32>; 6]),
     ([Value<i32>; 6], [Value<i32>; 6]),
     ([Value<i32>; 13], [Value<i32>; 13])) = (([Value::Some('y'),
                                                Value::Some('u'),
                                                Value::Some('a'),
                                                Value::Some('n'),
                                                Value::Some('c'),
                                                Value::Some('h'),
                                                Value::Some('a'),
                                                Value::Some('o')],
                                               [Value::Some('a'),
                                                Value::Some('a'),
                                                Value::Some('c'),
                                                Value::Some('h'),
                                                Value::Some('n'),
                                                Value::Some('o'),
                                                Value::Some('u'),
                                                Value::Some('y')]),
                                              ([Value::Some(5),
                                                Value::Some(2),
                                                Value::Some(4),
                                                Value::Some(6),
                                                Value::Some(1),
                                                Value::Some(3)],
                                               [Value::Some(1),
                                                Value::Some(2),
                                                Value::Some(3),
                                                Value::Some(4),
                                                Value::Some(5),
                                                Value::Some(6)]),
                                              ([], []),
                                              ([Value::Some(1.0)], [Value::Some(1.0)]),
                                              ([Value::Some(31),
                                                Value::Some(41),
                                                Value::Some(59),
                                                Value::Some(26),
                                                Value::Some(41),
                                                Value::Some(58)],
                                               [Value::Some(26),
                                                Value::Some(31),
                                                Value::Some(41),
                                                Value::Some(41),
                                                Value::Some(58),
                                                Value::Some(59)]),
                                              ([Value::Some(1),
                                                Value::Some(2),
                                                Value::Some(3),
                                                Value::Some(4),
                                                Value::Some(5),
                                                Value::Some(6)],
                                               [Value::Some(1),
                                                Value::Some(2),
                                                Value::Some(3),
                                                Value::Some(4),
                                                Value::Some(5),
                                                Value::Some(6)]),
                                              ([Value::Some(6),
                                                Value::Some(5),
                                                Value::Some(4),
                                                Value::Some(3),
                                                Value::Some(2),
                                                Value::Some(1)],
                                               [Value::Some(1),
                                                Value::Some(2),
                                                Value::Some(3),
                                                Value::Some(4),
                                                Value::Some(5),
                                                Value::Some(6)]),
                                              ([Value::Some(74),
                                                Value::Some(59),
                                                Value::Some(238),
                                                Value::Some(-784),
                                                Value::Some(9845),
                                                Value::Some(959),
                                                Value::Some(905),
                                                Value::Some(0),
                                                Value::Some(0),
                                                Value::Some(42),
                                                Value::Some(7586),
                                                Value::Some(-5467984),
                                                Value::Some(7586)],
                                               [Value::Some(-5467984),
                                                Value::Some(-784),
                                                Value::Some(0),
                                                Value::Some(0),
                                                Value::Some(42),
                                                Value::Some(59),
                                                Value::Some(74),
                                                Value::Some(238),
                                                Value::Some(905),
                                                Value::Some(959),
                                                Value::Some(7586),
                                                Value::Some(7586),
                                                Value::Some(9845)]));
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

pub mod ch01;
pub mod ch02;
