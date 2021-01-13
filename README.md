# cutlister

Generate the most efficient cutlist - don't waste your precious materials.

## Usage

Assume the following lengths and numbers of needed parts from 1000 mm stock.

| length | count |
| ------ | ----- |
| 200 mm | 8     |
| 280 mm | 2     |
| 540 mm | 3     |

Enter the command as follows:

```
cutlister 1000 3x200 2x280 1x540
```

The output in this example will be:

```
stock length: 1000
6 parts
cut list:
stock 1
  part 1    200
  part 2    200
  part 3    280
  part 4    280

  offcut    40

stock 2
  part 1    200
  part 2    540

  offcut    260

total offcut: 40
```

The parts are cramped into as few stock parts as possible, minimizing offcut. The remainder on the last stock can often be used for the next project and doesn't count towards the total offcut.

Note: The saw kerf width needs to be accounted for manually. When working with wood, stock dimensions aren't usually very precise, so removing a few mm from the actual stock length is advisable.
