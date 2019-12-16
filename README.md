# comment

A golang implementation of a commenting system, similar to Reddit

## Description
Comments are made of the following key elements:
- text
- parent
- siblings
- children
- score

Users need to be able to create new comments, list comments (a subtree) and vote
on comments.

Although it is not supported here, comments could also be deleted. On Reddit,
this manifests by preserving the tree structure but editing a comments text to
`[deleted]`.
