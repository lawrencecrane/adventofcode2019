use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;
use std::collections::HashMap;

fn main() {
    let map = input(&parse);

    println!("Total number of direct and indirect orbits {}", orbit_count(&map));
    println!("Minimum number of orbital transfers {}",
             find_shortest(&map, &String::from("YOU"), &String::from("SAN")));
}

fn find_shortest(map: &HashMap<String, Vec<String>>, a: &String, b: &String) -> usize {
    let mut steps = HashMap::new();

    let mut a_parents = map.get(a).unwrap().to_owned();
    let mut b_parents = map.get(b).unwrap().to_owned();

    for step in 0.. {
        match add_step_count(&mut steps, &a_parents, &b_parents, step) {
            Some(count) => return count,
            _ => ()
        }

        a_parents = combine_parents(a_parents, map);
        b_parents = combine_parents(b_parents, map);
    }

    return 0
}

fn add_step_count(steps: &mut HashMap<String, usize>,
                  nodes_a: &Vec<String>,
                  nodes_b: &Vec<String>,
                  step: usize) -> Option<usize> {
    let (_, final_count) = nodes_a.iter()
        .chain(nodes_b.iter())
        .fold((steps, None), |(steps, n), node| {
            if steps.contains_key(node) {
                let count = steps.get(node).unwrap() + step;
                return (steps, Some(count))
            }

            steps.insert(node.to_owned(), step);

            (steps, n)
        });

    return final_count
}

fn combine_parents(childs: Vec<String>, map: &HashMap<String, Vec<String>>) -> Vec<String> {
    childs.iter().fold(Vec::new(), |mut parents, child| {
        parents.extend_from_slice(map.get(child).unwrap());
        parents
    })
}

fn orbit_count(map: &HashMap<String, Vec<String>>) -> usize {
    let counts = direct_orbit_counts(&map);
    direct_orbit_count(&counts) + indirect_orbit_total_count(map, &counts)
}

fn indirect_orbit_total_count(map: &HashMap<String, Vec<String>>,
                              counts: &HashMap<&String, usize>) -> usize {
    map.keys().map(|val| indirect_orbit_count(val, map, counts)).sum()
}

fn indirect_orbit_count(val: &String,
                        map: &HashMap<String, Vec<String>>,
                        counts: &HashMap<&String, usize>) -> usize {
   let parents = map.get(val).unwrap();

   let sum: usize = parents.iter().map(|x| counts.get(x).unwrap()).sum();
   let sum2: usize = parents.iter().map(|x| indirect_orbit_count(x, map, counts)).sum();

   sum + sum2
}

fn direct_orbit_count(counts: &HashMap<&String, usize>) -> usize {
    counts.values().sum()
}

fn direct_orbit_counts(map: &HashMap<String, Vec<String>>) -> HashMap<&String, usize> {
    map.iter()
        .fold(HashMap::new(), |mut map, (key, parents)| {
            map.insert(key, parents.len());
            map
        })
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
    fn test_find_shortest() {
        let map = test_data2();
        assert_eq!(find_shortest(&map, &String::from("YOU"), &String::from("SAN")), 4)
    }

    #[test]
    fn test_indirect_orbit_count() {
        let map = test_data();
        let counts = direct_orbit_counts(&map);
        assert_eq!(indirect_orbit_total_count(&map, &counts), 31)
    }

    #[test]
    fn test_direct_orbit_count() {
        let map = test_data();
        let counts = direct_orbit_counts(&map);
        assert_eq!(direct_orbit_count(&counts), 11)
    }

    fn test_data2() -> HashMap<String, Vec<String>> {
        let data = ["COM)B", "B)C", "C)D", "D)E", "E)F",
            "B)G", "G)H", "D)I", "E)J", "J)K","K)L", "K)YOU", "I)SAN"];

        let mut map = HashMap::new();

        for val in data.iter() {
            parse(String::from(*val), &mut map);
        }

        map
    }

    fn test_data() -> HashMap<String, Vec<String>> {
        let data = ["COM)B", "B)C", "C)D", "D)E", "E)F",
            "B)G", "G)H", "D)I", "E)J", "J)K","K)L"];

        let mut map = HashMap::new();

        for val in data.iter() {
            parse(String::from(*val), &mut map);
        }

        map
    }
}
