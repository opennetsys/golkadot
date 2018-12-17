package triecodec

const EMPTY_TRIE = 0
const LEAF_NODE_OFFSET = 1
const LEAF_NODE_BIG = 127
const EXTENSION_NODE_OFFSET = 128
const EXTENSION_NODE_BIG = 253
const BRANCH_NODE_NO_VALUE = 254
const BRANCH_NODE_WITH_VALUE = 255
const LEAF_NODE_THRESHOLD = LEAF_NODE_BIG - LEAF_NODE_OFFSET
const EXTENSION_NODE_THRESHOLD = EXTENSION_NODE_BIG - EXTENSION_NODE_OFFSET // 125
const LEAF_NODE_SMALL_MAX = LEAF_NODE_BIG - 1
const EXTENSION_NODE_SMALL_MAX = EXTENSION_NODE_BIG - 1
const NODE_TYPE_NULL = 0
const NODE_TYPE_BRANCH = 1
const NODE_TYPE_EXT = 2
const NODE_TYPE_LEAF = 3
