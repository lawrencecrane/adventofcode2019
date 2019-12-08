use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

type Layer = Vec<Matrix>;
type Matrix = Vec<Vec<usize>>;

fn main() {
    println!("{:?}", to_layers(input(), 25, 6))
}

fn to_layers(mut data: String, width: usize, height: usize) -> Layer {
    let size = width * height;

    (0..data.len() / size)
        .map(|_| to_matrix(data.drain(0..size).collect(), width, height))
        .collect()
}

fn to_matrix(mut data: String, width: usize, height: usize) -> Matrix {
    (0..height)
        .map(|_| data.drain(0..width)
             .map(|x| x.to_digit(10).unwrap() as usize).collect())
        .collect()
}

fn input() -> String {
    let f = File::open("data/input").unwrap();
    let f = BufReader::new(f);

    f.lines().next().unwrap().unwrap()
}

// #[cfg(test)]
// mod tests {
//     use super::*;

//     #[test]
//     fn test_fn() {
//     }
// }
