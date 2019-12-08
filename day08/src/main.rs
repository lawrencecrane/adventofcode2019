use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

type Matrix = Vec<Vec<usize>>;

fn main() {
    let layers = to_layers(input(), 25, 6);
    let layer = find_layer_with_fewest_zeros(&layers);
    println!("Answer to Part 1: {}", multiply_counts_of(1, 2, layer))
}

fn multiply_counts_of(a: usize, b: usize, layer: &Vec<usize>) -> usize {
    count_of(a, layer) * count_of(b, layer)
}

fn find_layer_with_fewest_zeros(layers: &Matrix) -> &Vec<usize> {
    let (_, layer) = layers.iter()
        .fold((None, None), |(nzero, lyr), layer| {
            let count = count_of(0, layer);

            match nzero {
                Some(n) if n < count => (nzero, lyr),
                _ => (Some(count), Some(layer))
            }
        });

    layer.unwrap()
}

fn count_of(a: usize, layer: &Vec<usize>) -> usize {
    layer.iter().filter(|x| x == &&a).count()
}

fn to_layers(mut data: String, width: usize, height: usize) -> Matrix {
    let size = width * height;

    (0..data.len() / size)
        .map(|_| data.drain(0..size)
             .map(|x| x.to_digit(10).unwrap() as usize).collect())
        .collect()
}

fn input() -> String {
    let f = File::open("data/input").unwrap();
    let f = BufReader::new(f);

    f.lines().next().unwrap().unwrap()
}
