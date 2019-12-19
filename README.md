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

The comments have a couple of rules for being displayed. If a comment has a
certain threshold of votes, it and all of its parents should be displayed.
Otherwise, the user must manually 'un-fold' the comment. Also, sibling comments
should be sorted by votes.
