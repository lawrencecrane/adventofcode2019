use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;
use std::collections::HashMap;

fn main() {
    let map = input(&parse);
    println!("{:?}", map);
}

fn direct_orbit_count(map: &HashMap<String, Vec<String>>) -> usize {
    map.iter()
        .map(|(_, val)| val.len())
        .sum()
}

fn add_to_orbit_map(map: &mut HashMap<String, Vec<String>>,
                    child: &str,
                    parent: &str) {
    let parents = map.entry(String::from(child)).or_insert(Vec::new());
    parents.push(String::from(parent));

    map.entry(String::from(parent)).or_insert(Vec::new());
}

fn parse(data: String, out: &mut HashMap<String, Vec<String>>) {
    let mut splitted = data.split(')');

    match (splitted.next(), splitted.next()) {
        (Some(parent), Some(child)) => {
            add_to_orbit_map(out, child, parent);
        },
        _ => ()
    }
}

fn input(parser: &dyn Fn(String, &mut HashMap<String, Vec<String>>))
    -> HashMap<String, Vec<String>> {
    let f = File::open("data/input").unwrap();
    let f = BufReader::new(f);

    let mut map = HashMap::new();

    for line in f.lines() {
        parser(line.unwrap(), &mut map);
    }

    map
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_direct_orbit_count() {
        let data = ["COM)B", "B)C", "C)D", "D)E", "E)F",
            "B)G", "G)H", "D)I", "E)J", "J)K","K)L"];

        let mut map = HashMap::new();

        for val in data.iter() {
            parse(String::from(*val), &mut map);
        }

        assert_eq!(direct_orbit_count(&map), 11)
    }
}
