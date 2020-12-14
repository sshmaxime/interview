1) In the illustration, let’s imagine I know the whole Merkle tree. Someone gives me L2 data block but I don’t trust him. How can I check if L2 data is valid?
-> If I know the whole tree I would need to get the sibling of Hash 0-1(L2), which is Hash 0-0(L1), then concatenate them together in order to produce a hash. Then compare this hash with Hash 0. If same then L2 data received is valid.

2) I know only the L3 block and the Merkle root. What minimum information do I need to check that the L3 block and the Merkle root belong to the same Merkle tree ?
-> I need the sibling of Hash 1-0 (L3), which is Hash 1-1 (L4) and their parent's sibling Hash 0.

3) What are some Merkle tree use cases?
-> In systems where data integrity needs to be done: Distributed systems most of all but not only, ie: database replication.