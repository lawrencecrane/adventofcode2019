use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

type Matrix = Vec<Vec<usize>>;

fn main() {
    let size = 25 * 6;
    let layers = to_layers(input(), size);

    println!("Answer to Part 1: {}",
             multiply_counts_of(1, 2, find_layer_with_fewest_zeros(&layers)));

    let image = to_image(&layers, size);

    println!("Decoded image:");
    print_image(&image, 25, 6);
}

fn print_image(image: &Vec<usize>, width: usize, height: usize) {
    for h in 0..height {
        let row = image[h*width..(h*width + width)].iter()
                 .map(|x| {
                     match x {
                         0 => ' ',
                         _ => 'X'
                     }
                 });

        for pixel in row {
            print!("{}", pixel);
        }

        println!();
    }
}

fn to_image(layers: &Matrix, size: usize) -> Vec<usize> {
    (0..size)
        .fold(Vec::new(), |mut img, pixel| {
            let pxl = layers.iter()
                .map(|layer| layer[pixel])
                .find(|pxl| pxl != &2)
                .unwrap();

            img.push(pxl);

            img
        })
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

fn to_layers(mut data: String, size: usize) -> Matrix {
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
