// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar CQLLexer;

options {
    caseInsensitive = true;
}

// Operators and Punctuators

LR_BRACKET   : '(';
RR_BRACKET   : ')';
LC_BRACKET   : '{';
RC_BRACKET   : '}';
LS_BRACKET   : '[';
RS_BRACKET   : ']';
COMMA        : ',';
SEMI         : ';';
COLON        : ':';
DOT          : '.';
STAR         : '*';
DIVIDE       : '/';
MODULE       : '%';
PLUS         : '+';
MINUSMINUS   : '--';
MINUS        : '-';
DQUOTE       : '"';
SQUOTE       : '\'';
OPERATOR_EQ  : '=';
OPERATOR_LT  : '<';
OPERATOR_GT  : '>';
OPERATOR_LTE : '<=';
OPERATOR_GTE : '>=';

// Keywords

K_ADD            : 'ADD';
K_AGGREGATE      : 'AGGREGATE';
K_ALL            : 'ALL';
K_ALLOW          : 'ALLOW';
K_ALTER          : 'ALTER';
K_AND            : 'AND';
K_ANY            : 'ANY';
K_APPLY          : 'APPLY';
K_AS             : 'AS';
K_ASC            : 'ASC';
K_AUTHORIZE      : 'AUTHORIZE';
K_BATCH          : 'BATCH';
K_BEGIN          : 'BEGIN';
K_BY             : 'BY';
K_CALLED         : 'CALLED';
K_CLUSTERING     : 'CLUSTERING';
K_COLUMNFAMILY   : 'COLUMNFAMILY';
K_COMPACT        : 'COMPACT';
K_CONSISTENCY    : 'CONSISTENCY';
K_CONTAINS       : 'CONTAINS';
K_CREATE         : 'CREATE';
K_CUSTOM         : 'CUSTOM';
K_DELETE         : 'DELETE';
K_DESC           : 'DESC';
K_DESCRIBE       : 'DESCRIBE';
K_DISTINCT       : 'DISTINCT';
K_DROP           : 'DROP';
K_DURABLE_WRITES : 'DURABLE_WRITES';
K_EACH_QUORUM    : 'EACH_QUORUM';
K_ENTRIES        : 'ENTRIES';
K_EXECUTE        : 'EXECUTE';
K_EXISTS         : 'EXISTS';
K_FALSE          : 'FALSE';
K_FILTERING      : 'FILTERING';
K_FINALFUNC      : 'FINALFUNC';
K_FROM           : 'FROM';
K_FULL           : 'FULL';
K_FUNCTION       : 'FUNCTION';
K_FUNCTIONS      : 'FUNCTIONS';
K_GRANT          : 'GRANT';
K_IF             : 'IF';
K_IN             : 'IN';
K_INDEX          : 'INDEX';
K_INFINITY       : 'INFINITY';
K_INITCOND       : 'INITCOND';
K_INPUT          : 'INPUT';
K_INSERT         : 'INSERT';
K_INTO           : 'INTO';
K_IS             : 'IS';
K_JSON           : 'JSON';
K_KEY            : 'KEY';
K_KEYS           : 'KEYS';
K_KEYSPACE       : 'KEYSPACE';
K_KEYSPACES      : 'KEYSPACES';
K_LANGUAGE       : 'LANGUAGE';
K_LEVEL          : 'LEVEL';
K_LIMIT          : 'LIMIT';
K_LOCAL_ONE      : 'LOCAL_ONE';
K_LOCAL_QUORUM   : 'LOCAL_QUORUM';
K_LOGGED         : 'LOGGED';
K_LOGIN          : 'LOGIN';
K_MATERIALIZED   : 'MATERIALIZED';
K_MODIFY         : 'MODIFY';
K_NAN            : 'NAN';
K_NORECURSIVE    : 'NORECURSIVE';
K_NOSUPERUSER    : 'NOSUPERUSER';
K_NOT            : 'NOT';
K_NULL           : 'NULL';
K_OF             : 'OF';
K_ON             : 'ON';
K_ONE            : 'ONE';
K_OPTIONS        : 'OPTIONS';
K_OR             : 'OR';
K_ORDER          : 'ORDER';
K_PARTITION      : 'PARTITION';
K_PASSWORD       : 'PASSWORD';
K_PER            : 'PER';
K_PERMISSION     : 'PERMISSION';
K_PERMISSIONS    : 'PERMISSIONS';
K_PLACEHOLDER    : '?';
K_PRIMARY        : 'PRIMARY';
K_QUORUM         : 'QUORUM';
K_RENAME         : 'RENAME';
K_REPLACE        : 'REPLACE';
K_REPLICATION    : 'REPLICATION';
K_RETURNS        : 'RETURNS';
K_REVOKE         : 'REVOKE';
K_ROLE           : 'ROLE';
K_ROLES          : 'ROLES';
K_SCHEMA         : 'SCHEMA';
K_SELECT         : 'SELECT';
K_SET            : 'SET';
K_SFUNC          : 'SFUNC';
K_STATIC         : 'STATIC';
K_STORAGE        : 'STORAGE';
K_STYPE          : 'STYPE';
K_SUPERUSER      : 'SUPERUSER';
K_TABLE          : 'TABLE';
K_THREE          : 'THREE';
K_TIMESTAMP      : 'TIMESTAMP';
K_TO             : 'TO';
K_TOKEN          : 'TOKEN';
K_TRIGGER        : 'TRIGGER';
K_TRUE           : 'TRUE';
K_TRUNCATE       : 'TRUNCATE';
K_TTL            : 'TTL';
K_TWO            : 'TWO';
K_TYPE           : 'TYPE';
K_UNLOGGED       : 'UNLOGGED';
K_UPDATE         : 'UPDATE';
K_USE            : 'USE';
K_USER           : 'USER';
K_USING          : 'USING';
K_UUID           : 'UUID';
K_VALUES         : 'VALUES';
K_VIEW           : 'VIEW';
K_WHERE          : 'WHERE';
K_WITH           : 'WITH';
K_WRITETIME      : 'WRITETIME';
K_ASCII          : 'ASCII';
K_BIGINT         : 'BIGINT';
K_BLOB           : 'BLOB';
K_BOOLEAN        : 'BOOLEAN';
K_COUNTER        : 'COUNTER';
K_DATE           : 'DATE';
K_DECIMAL        : 'DECIMAL';
K_DOUBLE         : 'DOUBLE';
K_FLOAT          : 'FLOAT';
K_FROZEN         : 'FROZEN';
K_INET           : 'INET';
K_INT            : 'INT';
K_LIST           : 'LIST';
K_MAP            : 'MAP';
K_SMALLINT       : 'SMALLINT';
K_TEXT           : 'TEXT';
K_TIMEUUID       : 'TIMEUUID';
K_TIME           : 'TIME';
K_TINYINT        : 'TINYINT';
K_TUPLE          : 'TUPLE';
K_VARCHAR        : 'VARCHAR';
K_VARINT         : 'VARINT';

// Literals

CODE_BLOCK: '$$' (~ '$' | '$' ~'$')* '$$';

STRING_LITERAL: '\'' ('\\' . | '\'\'' | ~('\'' | '\\'))* '\'';

DECIMAL_LITERAL: DEC_DIGIT+;

FLOAT_LITERAL: MINUS? [0-9]+ (DOT [0-9]+)?;

HEXADECIMAL_LITERAL: 'X' '\'' (HEX_DIGIT HEX_DIGIT)+ '\'' | '0X' HEX_DIGIT+;

REAL_LITERAL: DEC_DIGIT+ '.'? EXPONENT_NUM_PART | DEC_DIGIT* '.' DEC_DIGIT+ EXPONENT_NUM_PART?;

OBJECT_NAME: [A-Z] [A-Z0-9_$]* | '"' ~'"'+ '"';

UUID:
    HEX_4DIGIT HEX_4DIGIT '-' HEX_4DIGIT '-' HEX_4DIGIT '-' HEX_4DIGIT '-' HEX_4DIGIT HEX_4DIGIT HEX_4DIGIT
;

// Hidden

SPACE              : [ \t\r\n]+     -> channel (HIDDEN);
SPEC_MYSQL_COMMENT : '/*!' .+? '*/' -> channel (HIDDEN);
COMMENT_INPUT      : '/*' .*? '*/'  -> channel (HIDDEN);
LINE_COMMENT:
    (('-- ' | '#' | '//') ~ [\r\n]* ('\r'? '\n' | EOF) | '--' ('\r'? '\n' | EOF)) -> channel (HIDDEN)
;

// Fragments

fragment HEX_4DIGIT: [0-9A-F] [0-9A-F] [0-9A-F] [0-9A-F];

fragment HEX_DIGIT: [0-9A-F];

fragment DEC_DIGIT: [0-9];

fragment EXPONENT_NUM_PART: 'E' '-'? DEC_DIGIT+;
