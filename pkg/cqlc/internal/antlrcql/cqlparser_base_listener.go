// Code generated from CQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlrcql // CQLParser
import "github.com/antlr4-go/antlr/v4"

// BaseCQLParserListener is a complete listener for a parse tree produced by CQLParser.
type BaseCQLParserListener struct{}

var _ CQLParserListener = &BaseCQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseCQLParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseCQLParserListener) ExitRoot(ctx *RootContext) {}

// EnterCqls is called when production cqls is entered.
func (s *BaseCQLParserListener) EnterCqls(ctx *CqlsContext) {}

// ExitCqls is called when production cqls is exited.
func (s *BaseCQLParserListener) ExitCqls(ctx *CqlsContext) {}

// EnterStatementSeparator is called when production statementSeparator is entered.
func (s *BaseCQLParserListener) EnterStatementSeparator(ctx *StatementSeparatorContext) {}

// ExitStatementSeparator is called when production statementSeparator is exited.
func (s *BaseCQLParserListener) ExitStatementSeparator(ctx *StatementSeparatorContext) {}

// EnterEmpty_ is called when production empty_ is entered.
func (s *BaseCQLParserListener) EnterEmpty_(ctx *Empty_Context) {}

// ExitEmpty_ is called when production empty_ is exited.
func (s *BaseCQLParserListener) ExitEmpty_(ctx *Empty_Context) {}

// EnterCql is called when production cql is entered.
func (s *BaseCQLParserListener) EnterCql(ctx *CqlContext) {}

// ExitCql is called when production cql is exited.
func (s *BaseCQLParserListener) ExitCql(ctx *CqlContext) {}

// EnterRevoke is called when production revoke is entered.
func (s *BaseCQLParserListener) EnterRevoke(ctx *RevokeContext) {}

// ExitRevoke is called when production revoke is exited.
func (s *BaseCQLParserListener) ExitRevoke(ctx *RevokeContext) {}

// EnterListRoles is called when production listRoles is entered.
func (s *BaseCQLParserListener) EnterListRoles(ctx *ListRolesContext) {}

// ExitListRoles is called when production listRoles is exited.
func (s *BaseCQLParserListener) ExitListRoles(ctx *ListRolesContext) {}

// EnterListPermissions is called when production listPermissions is entered.
func (s *BaseCQLParserListener) EnterListPermissions(ctx *ListPermissionsContext) {}

// ExitListPermissions is called when production listPermissions is exited.
func (s *BaseCQLParserListener) ExitListPermissions(ctx *ListPermissionsContext) {}

// EnterGrant is called when production grant is entered.
func (s *BaseCQLParserListener) EnterGrant(ctx *GrantContext) {}

// ExitGrant is called when production grant is exited.
func (s *BaseCQLParserListener) ExitGrant(ctx *GrantContext) {}

// EnterPriviledge is called when production priviledge is entered.
func (s *BaseCQLParserListener) EnterPriviledge(ctx *PriviledgeContext) {}

// ExitPriviledge is called when production priviledge is exited.
func (s *BaseCQLParserListener) ExitPriviledge(ctx *PriviledgeContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseCQLParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseCQLParserListener) ExitResource(ctx *ResourceContext) {}

// EnterCreateUser is called when production createUser is entered.
func (s *BaseCQLParserListener) EnterCreateUser(ctx *CreateUserContext) {}

// ExitCreateUser is called when production createUser is exited.
func (s *BaseCQLParserListener) ExitCreateUser(ctx *CreateUserContext) {}

// EnterCreateRole is called when production createRole is entered.
func (s *BaseCQLParserListener) EnterCreateRole(ctx *CreateRoleContext) {}

// ExitCreateRole is called when production createRole is exited.
func (s *BaseCQLParserListener) ExitCreateRole(ctx *CreateRoleContext) {}

// EnterCreateType is called when production createType is entered.
func (s *BaseCQLParserListener) EnterCreateType(ctx *CreateTypeContext) {}

// ExitCreateType is called when production createType is exited.
func (s *BaseCQLParserListener) ExitCreateType(ctx *CreateTypeContext) {}

// EnterTypeMemberColumnList is called when production typeMemberColumnList is entered.
func (s *BaseCQLParserListener) EnterTypeMemberColumnList(ctx *TypeMemberColumnListContext) {}

// ExitTypeMemberColumnList is called when production typeMemberColumnList is exited.
func (s *BaseCQLParserListener) ExitTypeMemberColumnList(ctx *TypeMemberColumnListContext) {}

// EnterCreateTrigger is called when production createTrigger is entered.
func (s *BaseCQLParserListener) EnterCreateTrigger(ctx *CreateTriggerContext) {}

// ExitCreateTrigger is called when production createTrigger is exited.
func (s *BaseCQLParserListener) ExitCreateTrigger(ctx *CreateTriggerContext) {}

// EnterCreateMaterializedView is called when production createMaterializedView is entered.
func (s *BaseCQLParserListener) EnterCreateMaterializedView(ctx *CreateMaterializedViewContext) {}

// ExitCreateMaterializedView is called when production createMaterializedView is exited.
func (s *BaseCQLParserListener) ExitCreateMaterializedView(ctx *CreateMaterializedViewContext) {}

// EnterMaterializedViewWhere is called when production materializedViewWhere is entered.
func (s *BaseCQLParserListener) EnterMaterializedViewWhere(ctx *MaterializedViewWhereContext) {}

// ExitMaterializedViewWhere is called when production materializedViewWhere is exited.
func (s *BaseCQLParserListener) ExitMaterializedViewWhere(ctx *MaterializedViewWhereContext) {}

// EnterColumnNotNullList is called when production columnNotNullList is entered.
func (s *BaseCQLParserListener) EnterColumnNotNullList(ctx *ColumnNotNullListContext) {}

// ExitColumnNotNullList is called when production columnNotNullList is exited.
func (s *BaseCQLParserListener) ExitColumnNotNullList(ctx *ColumnNotNullListContext) {}

// EnterColumnNotNull is called when production columnNotNull is entered.
func (s *BaseCQLParserListener) EnterColumnNotNull(ctx *ColumnNotNullContext) {}

// ExitColumnNotNull is called when production columnNotNull is exited.
func (s *BaseCQLParserListener) ExitColumnNotNull(ctx *ColumnNotNullContext) {}

// EnterMaterializedViewOptions is called when production materializedViewOptions is entered.
func (s *BaseCQLParserListener) EnterMaterializedViewOptions(ctx *MaterializedViewOptionsContext) {}

// ExitMaterializedViewOptions is called when production materializedViewOptions is exited.
func (s *BaseCQLParserListener) ExitMaterializedViewOptions(ctx *MaterializedViewOptionsContext) {}

// EnterCreateKeyspace is called when production createKeyspace is entered.
func (s *BaseCQLParserListener) EnterCreateKeyspace(ctx *CreateKeyspaceContext) {}

// ExitCreateKeyspace is called when production createKeyspace is exited.
func (s *BaseCQLParserListener) ExitCreateKeyspace(ctx *CreateKeyspaceContext) {}

// EnterCreateFunction is called when production createFunction is entered.
func (s *BaseCQLParserListener) EnterCreateFunction(ctx *CreateFunctionContext) {}

// ExitCreateFunction is called when production createFunction is exited.
func (s *BaseCQLParserListener) ExitCreateFunction(ctx *CreateFunctionContext) {}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *BaseCQLParserListener) EnterCodeBlock(ctx *CodeBlockContext) {}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *BaseCQLParserListener) ExitCodeBlock(ctx *CodeBlockContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseCQLParserListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseCQLParserListener) ExitParamList(ctx *ParamListContext) {}

// EnterReturnMode is called when production returnMode is entered.
func (s *BaseCQLParserListener) EnterReturnMode(ctx *ReturnModeContext) {}

// ExitReturnMode is called when production returnMode is exited.
func (s *BaseCQLParserListener) ExitReturnMode(ctx *ReturnModeContext) {}

// EnterCreateAggregate is called when production createAggregate is entered.
func (s *BaseCQLParserListener) EnterCreateAggregate(ctx *CreateAggregateContext) {}

// ExitCreateAggregate is called when production createAggregate is exited.
func (s *BaseCQLParserListener) ExitCreateAggregate(ctx *CreateAggregateContext) {}

// EnterInitCondDefinition is called when production initCondDefinition is entered.
func (s *BaseCQLParserListener) EnterInitCondDefinition(ctx *InitCondDefinitionContext) {}

// ExitInitCondDefinition is called when production initCondDefinition is exited.
func (s *BaseCQLParserListener) ExitInitCondDefinition(ctx *InitCondDefinitionContext) {}

// EnterInitCondHash is called when production initCondHash is entered.
func (s *BaseCQLParserListener) EnterInitCondHash(ctx *InitCondHashContext) {}

// ExitInitCondHash is called when production initCondHash is exited.
func (s *BaseCQLParserListener) ExitInitCondHash(ctx *InitCondHashContext) {}

// EnterInitCondHashItem is called when production initCondHashItem is entered.
func (s *BaseCQLParserListener) EnterInitCondHashItem(ctx *InitCondHashItemContext) {}

// ExitInitCondHashItem is called when production initCondHashItem is exited.
func (s *BaseCQLParserListener) ExitInitCondHashItem(ctx *InitCondHashItemContext) {}

// EnterInitCondListNested is called when production initCondListNested is entered.
func (s *BaseCQLParserListener) EnterInitCondListNested(ctx *InitCondListNestedContext) {}

// ExitInitCondListNested is called when production initCondListNested is exited.
func (s *BaseCQLParserListener) ExitInitCondListNested(ctx *InitCondListNestedContext) {}

// EnterInitCondList is called when production initCondList is entered.
func (s *BaseCQLParserListener) EnterInitCondList(ctx *InitCondListContext) {}

// ExitInitCondList is called when production initCondList is exited.
func (s *BaseCQLParserListener) ExitInitCondList(ctx *InitCondListContext) {}

// EnterOrReplace is called when production orReplace is entered.
func (s *BaseCQLParserListener) EnterOrReplace(ctx *OrReplaceContext) {}

// ExitOrReplace is called when production orReplace is exited.
func (s *BaseCQLParserListener) ExitOrReplace(ctx *OrReplaceContext) {}

// EnterAlterUser is called when production alterUser is entered.
func (s *BaseCQLParserListener) EnterAlterUser(ctx *AlterUserContext) {}

// ExitAlterUser is called when production alterUser is exited.
func (s *BaseCQLParserListener) ExitAlterUser(ctx *AlterUserContext) {}

// EnterUserPassword is called when production userPassword is entered.
func (s *BaseCQLParserListener) EnterUserPassword(ctx *UserPasswordContext) {}

// ExitUserPassword is called when production userPassword is exited.
func (s *BaseCQLParserListener) ExitUserPassword(ctx *UserPasswordContext) {}

// EnterUserSuperUser is called when production userSuperUser is entered.
func (s *BaseCQLParserListener) EnterUserSuperUser(ctx *UserSuperUserContext) {}

// ExitUserSuperUser is called when production userSuperUser is exited.
func (s *BaseCQLParserListener) ExitUserSuperUser(ctx *UserSuperUserContext) {}

// EnterAlterType is called when production alterType is entered.
func (s *BaseCQLParserListener) EnterAlterType(ctx *AlterTypeContext) {}

// ExitAlterType is called when production alterType is exited.
func (s *BaseCQLParserListener) ExitAlterType(ctx *AlterTypeContext) {}

// EnterAlterTypeOperation is called when production alterTypeOperation is entered.
func (s *BaseCQLParserListener) EnterAlterTypeOperation(ctx *AlterTypeOperationContext) {}

// ExitAlterTypeOperation is called when production alterTypeOperation is exited.
func (s *BaseCQLParserListener) ExitAlterTypeOperation(ctx *AlterTypeOperationContext) {}

// EnterAlterTypeRename is called when production alterTypeRename is entered.
func (s *BaseCQLParserListener) EnterAlterTypeRename(ctx *AlterTypeRenameContext) {}

// ExitAlterTypeRename is called when production alterTypeRename is exited.
func (s *BaseCQLParserListener) ExitAlterTypeRename(ctx *AlterTypeRenameContext) {}

// EnterAlterTypeRenameList is called when production alterTypeRenameList is entered.
func (s *BaseCQLParserListener) EnterAlterTypeRenameList(ctx *AlterTypeRenameListContext) {}

// ExitAlterTypeRenameList is called when production alterTypeRenameList is exited.
func (s *BaseCQLParserListener) ExitAlterTypeRenameList(ctx *AlterTypeRenameListContext) {}

// EnterAlterTypeRenameItem is called when production alterTypeRenameItem is entered.
func (s *BaseCQLParserListener) EnterAlterTypeRenameItem(ctx *AlterTypeRenameItemContext) {}

// ExitAlterTypeRenameItem is called when production alterTypeRenameItem is exited.
func (s *BaseCQLParserListener) ExitAlterTypeRenameItem(ctx *AlterTypeRenameItemContext) {}

// EnterAlterTypeAdd is called when production alterTypeAdd is entered.
func (s *BaseCQLParserListener) EnterAlterTypeAdd(ctx *AlterTypeAddContext) {}

// ExitAlterTypeAdd is called when production alterTypeAdd is exited.
func (s *BaseCQLParserListener) ExitAlterTypeAdd(ctx *AlterTypeAddContext) {}

// EnterAlterTypeAlterType is called when production alterTypeAlterType is entered.
func (s *BaseCQLParserListener) EnterAlterTypeAlterType(ctx *AlterTypeAlterTypeContext) {}

// ExitAlterTypeAlterType is called when production alterTypeAlterType is exited.
func (s *BaseCQLParserListener) ExitAlterTypeAlterType(ctx *AlterTypeAlterTypeContext) {}

// EnterAlterTable is called when production alterTable is entered.
func (s *BaseCQLParserListener) EnterAlterTable(ctx *AlterTableContext) {}

// ExitAlterTable is called when production alterTable is exited.
func (s *BaseCQLParserListener) ExitAlterTable(ctx *AlterTableContext) {}

// EnterAlterTableOperation is called when production alterTableOperation is entered.
func (s *BaseCQLParserListener) EnterAlterTableOperation(ctx *AlterTableOperationContext) {}

// ExitAlterTableOperation is called when production alterTableOperation is exited.
func (s *BaseCQLParserListener) ExitAlterTableOperation(ctx *AlterTableOperationContext) {}

// EnterAlterTableWith is called when production alterTableWith is entered.
func (s *BaseCQLParserListener) EnterAlterTableWith(ctx *AlterTableWithContext) {}

// ExitAlterTableWith is called when production alterTableWith is exited.
func (s *BaseCQLParserListener) ExitAlterTableWith(ctx *AlterTableWithContext) {}

// EnterAlterTableRename is called when production alterTableRename is entered.
func (s *BaseCQLParserListener) EnterAlterTableRename(ctx *AlterTableRenameContext) {}

// ExitAlterTableRename is called when production alterTableRename is exited.
func (s *BaseCQLParserListener) ExitAlterTableRename(ctx *AlterTableRenameContext) {}

// EnterAlterTableDropCompactStorage is called when production alterTableDropCompactStorage is entered.
func (s *BaseCQLParserListener) EnterAlterTableDropCompactStorage(ctx *AlterTableDropCompactStorageContext) {
}

// ExitAlterTableDropCompactStorage is called when production alterTableDropCompactStorage is exited.
func (s *BaseCQLParserListener) ExitAlterTableDropCompactStorage(ctx *AlterTableDropCompactStorageContext) {
}

// EnterAlterTableDropColumns is called when production alterTableDropColumns is entered.
func (s *BaseCQLParserListener) EnterAlterTableDropColumns(ctx *AlterTableDropColumnsContext) {}

// ExitAlterTableDropColumns is called when production alterTableDropColumns is exited.
func (s *BaseCQLParserListener) ExitAlterTableDropColumns(ctx *AlterTableDropColumnsContext) {}

// EnterAlterTableDropColumnList is called when production alterTableDropColumnList is entered.
func (s *BaseCQLParserListener) EnterAlterTableDropColumnList(ctx *AlterTableDropColumnListContext) {}

// ExitAlterTableDropColumnList is called when production alterTableDropColumnList is exited.
func (s *BaseCQLParserListener) ExitAlterTableDropColumnList(ctx *AlterTableDropColumnListContext) {}

// EnterAlterTableAdd is called when production alterTableAdd is entered.
func (s *BaseCQLParserListener) EnterAlterTableAdd(ctx *AlterTableAddContext) {}

// ExitAlterTableAdd is called when production alterTableAdd is exited.
func (s *BaseCQLParserListener) ExitAlterTableAdd(ctx *AlterTableAddContext) {}

// EnterAlterTableColumnDefinition is called when production alterTableColumnDefinition is entered.
func (s *BaseCQLParserListener) EnterAlterTableColumnDefinition(ctx *AlterTableColumnDefinitionContext) {
}

// ExitAlterTableColumnDefinition is called when production alterTableColumnDefinition is exited.
func (s *BaseCQLParserListener) ExitAlterTableColumnDefinition(ctx *AlterTableColumnDefinitionContext) {
}

// EnterAlterRole is called when production alterRole is entered.
func (s *BaseCQLParserListener) EnterAlterRole(ctx *AlterRoleContext) {}

// ExitAlterRole is called when production alterRole is exited.
func (s *BaseCQLParserListener) ExitAlterRole(ctx *AlterRoleContext) {}

// EnterRoleWith is called when production roleWith is entered.
func (s *BaseCQLParserListener) EnterRoleWith(ctx *RoleWithContext) {}

// ExitRoleWith is called when production roleWith is exited.
func (s *BaseCQLParserListener) ExitRoleWith(ctx *RoleWithContext) {}

// EnterRoleWithOptions is called when production roleWithOptions is entered.
func (s *BaseCQLParserListener) EnterRoleWithOptions(ctx *RoleWithOptionsContext) {}

// ExitRoleWithOptions is called when production roleWithOptions is exited.
func (s *BaseCQLParserListener) ExitRoleWithOptions(ctx *RoleWithOptionsContext) {}

// EnterAlterMaterializedView is called when production alterMaterializedView is entered.
func (s *BaseCQLParserListener) EnterAlterMaterializedView(ctx *AlterMaterializedViewContext) {}

// ExitAlterMaterializedView is called when production alterMaterializedView is exited.
func (s *BaseCQLParserListener) ExitAlterMaterializedView(ctx *AlterMaterializedViewContext) {}

// EnterDropUser is called when production dropUser is entered.
func (s *BaseCQLParserListener) EnterDropUser(ctx *DropUserContext) {}

// ExitDropUser is called when production dropUser is exited.
func (s *BaseCQLParserListener) ExitDropUser(ctx *DropUserContext) {}

// EnterDropType is called when production dropType is entered.
func (s *BaseCQLParserListener) EnterDropType(ctx *DropTypeContext) {}

// ExitDropType is called when production dropType is exited.
func (s *BaseCQLParserListener) ExitDropType(ctx *DropTypeContext) {}

// EnterDropMaterializedView is called when production dropMaterializedView is entered.
func (s *BaseCQLParserListener) EnterDropMaterializedView(ctx *DropMaterializedViewContext) {}

// ExitDropMaterializedView is called when production dropMaterializedView is exited.
func (s *BaseCQLParserListener) ExitDropMaterializedView(ctx *DropMaterializedViewContext) {}

// EnterDropAggregate is called when production dropAggregate is entered.
func (s *BaseCQLParserListener) EnterDropAggregate(ctx *DropAggregateContext) {}

// ExitDropAggregate is called when production dropAggregate is exited.
func (s *BaseCQLParserListener) ExitDropAggregate(ctx *DropAggregateContext) {}

// EnterDropFunction is called when production dropFunction is entered.
func (s *BaseCQLParserListener) EnterDropFunction(ctx *DropFunctionContext) {}

// ExitDropFunction is called when production dropFunction is exited.
func (s *BaseCQLParserListener) ExitDropFunction(ctx *DropFunctionContext) {}

// EnterDropTrigger is called when production dropTrigger is entered.
func (s *BaseCQLParserListener) EnterDropTrigger(ctx *DropTriggerContext) {}

// ExitDropTrigger is called when production dropTrigger is exited.
func (s *BaseCQLParserListener) ExitDropTrigger(ctx *DropTriggerContext) {}

// EnterDropRole is called when production dropRole is entered.
func (s *BaseCQLParserListener) EnterDropRole(ctx *DropRoleContext) {}

// ExitDropRole is called when production dropRole is exited.
func (s *BaseCQLParserListener) ExitDropRole(ctx *DropRoleContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseCQLParserListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseCQLParserListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropKeyspace is called when production dropKeyspace is entered.
func (s *BaseCQLParserListener) EnterDropKeyspace(ctx *DropKeyspaceContext) {}

// ExitDropKeyspace is called when production dropKeyspace is exited.
func (s *BaseCQLParserListener) ExitDropKeyspace(ctx *DropKeyspaceContext) {}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseCQLParserListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseCQLParserListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseCQLParserListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseCQLParserListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterWithElement is called when production withElement is entered.
func (s *BaseCQLParserListener) EnterWithElement(ctx *WithElementContext) {}

// ExitWithElement is called when production withElement is exited.
func (s *BaseCQLParserListener) ExitWithElement(ctx *WithElementContext) {}

// EnterTableOptions is called when production tableOptions is entered.
func (s *BaseCQLParserListener) EnterTableOptions(ctx *TableOptionsContext) {}

// ExitTableOptions is called when production tableOptions is exited.
func (s *BaseCQLParserListener) ExitTableOptions(ctx *TableOptionsContext) {}

// EnterClusteringOrder is called when production clusteringOrder is entered.
func (s *BaseCQLParserListener) EnterClusteringOrder(ctx *ClusteringOrderContext) {}

// ExitClusteringOrder is called when production clusteringOrder is exited.
func (s *BaseCQLParserListener) ExitClusteringOrder(ctx *ClusteringOrderContext) {}

// EnterTableOptionItem is called when production tableOptionItem is entered.
func (s *BaseCQLParserListener) EnterTableOptionItem(ctx *TableOptionItemContext) {}

// ExitTableOptionItem is called when production tableOptionItem is exited.
func (s *BaseCQLParserListener) ExitTableOptionItem(ctx *TableOptionItemContext) {}

// EnterTableOptionName is called when production tableOptionName is entered.
func (s *BaseCQLParserListener) EnterTableOptionName(ctx *TableOptionNameContext) {}

// ExitTableOptionName is called when production tableOptionName is exited.
func (s *BaseCQLParserListener) ExitTableOptionName(ctx *TableOptionNameContext) {}

// EnterTableOptionValue is called when production tableOptionValue is entered.
func (s *BaseCQLParserListener) EnterTableOptionValue(ctx *TableOptionValueContext) {}

// ExitTableOptionValue is called when production tableOptionValue is exited.
func (s *BaseCQLParserListener) ExitTableOptionValue(ctx *TableOptionValueContext) {}

// EnterOptionHash is called when production optionHash is entered.
func (s *BaseCQLParserListener) EnterOptionHash(ctx *OptionHashContext) {}

// ExitOptionHash is called when production optionHash is exited.
func (s *BaseCQLParserListener) ExitOptionHash(ctx *OptionHashContext) {}

// EnterOptionHashItem is called when production optionHashItem is entered.
func (s *BaseCQLParserListener) EnterOptionHashItem(ctx *OptionHashItemContext) {}

// ExitOptionHashItem is called when production optionHashItem is exited.
func (s *BaseCQLParserListener) ExitOptionHashItem(ctx *OptionHashItemContext) {}

// EnterOptionHashKey is called when production optionHashKey is entered.
func (s *BaseCQLParserListener) EnterOptionHashKey(ctx *OptionHashKeyContext) {}

// ExitOptionHashKey is called when production optionHashKey is exited.
func (s *BaseCQLParserListener) ExitOptionHashKey(ctx *OptionHashKeyContext) {}

// EnterOptionHashValue is called when production optionHashValue is entered.
func (s *BaseCQLParserListener) EnterOptionHashValue(ctx *OptionHashValueContext) {}

// ExitOptionHashValue is called when production optionHashValue is exited.
func (s *BaseCQLParserListener) ExitOptionHashValue(ctx *OptionHashValueContext) {}

// EnterColumnDefinitionList is called when production columnDefinitionList is entered.
func (s *BaseCQLParserListener) EnterColumnDefinitionList(ctx *ColumnDefinitionListContext) {}

// ExitColumnDefinitionList is called when production columnDefinitionList is exited.
func (s *BaseCQLParserListener) ExitColumnDefinitionList(ctx *ColumnDefinitionListContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseCQLParserListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseCQLParserListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterPrimaryKeyColumn is called when production primaryKeyColumn is entered.
func (s *BaseCQLParserListener) EnterPrimaryKeyColumn(ctx *PrimaryKeyColumnContext) {}

// ExitPrimaryKeyColumn is called when production primaryKeyColumn is exited.
func (s *BaseCQLParserListener) ExitPrimaryKeyColumn(ctx *PrimaryKeyColumnContext) {}

// EnterPrimaryKeyElement is called when production primaryKeyElement is entered.
func (s *BaseCQLParserListener) EnterPrimaryKeyElement(ctx *PrimaryKeyElementContext) {}

// ExitPrimaryKeyElement is called when production primaryKeyElement is exited.
func (s *BaseCQLParserListener) ExitPrimaryKeyElement(ctx *PrimaryKeyElementContext) {}

// EnterPrimaryKeyDefinition is called when production primaryKeyDefinition is entered.
func (s *BaseCQLParserListener) EnterPrimaryKeyDefinition(ctx *PrimaryKeyDefinitionContext) {}

// ExitPrimaryKeyDefinition is called when production primaryKeyDefinition is exited.
func (s *BaseCQLParserListener) ExitPrimaryKeyDefinition(ctx *PrimaryKeyDefinitionContext) {}

// EnterSinglePrimaryKey is called when production singlePrimaryKey is entered.
func (s *BaseCQLParserListener) EnterSinglePrimaryKey(ctx *SinglePrimaryKeyContext) {}

// ExitSinglePrimaryKey is called when production singlePrimaryKey is exited.
func (s *BaseCQLParserListener) ExitSinglePrimaryKey(ctx *SinglePrimaryKeyContext) {}

// EnterCompoundKey is called when production compoundKey is entered.
func (s *BaseCQLParserListener) EnterCompoundKey(ctx *CompoundKeyContext) {}

// ExitCompoundKey is called when production compoundKey is exited.
func (s *BaseCQLParserListener) ExitCompoundKey(ctx *CompoundKeyContext) {}

// EnterCompositeKey is called when production compositeKey is entered.
func (s *BaseCQLParserListener) EnterCompositeKey(ctx *CompositeKeyContext) {}

// ExitCompositeKey is called when production compositeKey is exited.
func (s *BaseCQLParserListener) ExitCompositeKey(ctx *CompositeKeyContext) {}

// EnterPartitionKeyList is called when production partitionKeyList is entered.
func (s *BaseCQLParserListener) EnterPartitionKeyList(ctx *PartitionKeyListContext) {}

// ExitPartitionKeyList is called when production partitionKeyList is exited.
func (s *BaseCQLParserListener) ExitPartitionKeyList(ctx *PartitionKeyListContext) {}

// EnterClusteringKeyList is called when production clusteringKeyList is entered.
func (s *BaseCQLParserListener) EnterClusteringKeyList(ctx *ClusteringKeyListContext) {}

// ExitClusteringKeyList is called when production clusteringKeyList is exited.
func (s *BaseCQLParserListener) ExitClusteringKeyList(ctx *ClusteringKeyListContext) {}

// EnterPartitionKey is called when production partitionKey is entered.
func (s *BaseCQLParserListener) EnterPartitionKey(ctx *PartitionKeyContext) {}

// ExitPartitionKey is called when production partitionKey is exited.
func (s *BaseCQLParserListener) ExitPartitionKey(ctx *PartitionKeyContext) {}

// EnterClusteringKey is called when production clusteringKey is entered.
func (s *BaseCQLParserListener) EnterClusteringKey(ctx *ClusteringKeyContext) {}

// ExitClusteringKey is called when production clusteringKey is exited.
func (s *BaseCQLParserListener) ExitClusteringKey(ctx *ClusteringKeyContext) {}

// EnterApplyBatch is called when production applyBatch is entered.
func (s *BaseCQLParserListener) EnterApplyBatch(ctx *ApplyBatchContext) {}

// ExitApplyBatch is called when production applyBatch is exited.
func (s *BaseCQLParserListener) ExitApplyBatch(ctx *ApplyBatchContext) {}

// EnterBeginBatch is called when production beginBatch is entered.
func (s *BaseCQLParserListener) EnterBeginBatch(ctx *BeginBatchContext) {}

// ExitBeginBatch is called when production beginBatch is exited.
func (s *BaseCQLParserListener) ExitBeginBatch(ctx *BeginBatchContext) {}

// EnterBatchType is called when production batchType is entered.
func (s *BaseCQLParserListener) EnterBatchType(ctx *BatchTypeContext) {}

// ExitBatchType is called when production batchType is exited.
func (s *BaseCQLParserListener) ExitBatchType(ctx *BatchTypeContext) {}

// EnterAlterKeyspace is called when production alterKeyspace is entered.
func (s *BaseCQLParserListener) EnterAlterKeyspace(ctx *AlterKeyspaceContext) {}

// ExitAlterKeyspace is called when production alterKeyspace is exited.
func (s *BaseCQLParserListener) ExitAlterKeyspace(ctx *AlterKeyspaceContext) {}

// EnterReplicationList is called when production replicationList is entered.
func (s *BaseCQLParserListener) EnterReplicationList(ctx *ReplicationListContext) {}

// ExitReplicationList is called when production replicationList is exited.
func (s *BaseCQLParserListener) ExitReplicationList(ctx *ReplicationListContext) {}

// EnterReplicationListItem is called when production replicationListItem is entered.
func (s *BaseCQLParserListener) EnterReplicationListItem(ctx *ReplicationListItemContext) {}

// ExitReplicationListItem is called when production replicationListItem is exited.
func (s *BaseCQLParserListener) ExitReplicationListItem(ctx *ReplicationListItemContext) {}

// EnterDurableWrites is called when production durableWrites is entered.
func (s *BaseCQLParserListener) EnterDurableWrites(ctx *DurableWritesContext) {}

// ExitDurableWrites is called when production durableWrites is exited.
func (s *BaseCQLParserListener) ExitDurableWrites(ctx *DurableWritesContext) {}

// EnterUse_ is called when production use_ is entered.
func (s *BaseCQLParserListener) EnterUse_(ctx *Use_Context) {}

// ExitUse_ is called when production use_ is exited.
func (s *BaseCQLParserListener) ExitUse_(ctx *Use_Context) {}

// EnterTruncate is called when production truncate is entered.
func (s *BaseCQLParserListener) EnterTruncate(ctx *TruncateContext) {}

// ExitTruncate is called when production truncate is exited.
func (s *BaseCQLParserListener) ExitTruncate(ctx *TruncateContext) {}

// EnterCreateIndex is called when production createIndex is entered.
func (s *BaseCQLParserListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production createIndex is exited.
func (s *BaseCQLParserListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseCQLParserListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseCQLParserListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterIndexColumnSpec is called when production indexColumnSpec is entered.
func (s *BaseCQLParserListener) EnterIndexColumnSpec(ctx *IndexColumnSpecContext) {}

// ExitIndexColumnSpec is called when production indexColumnSpec is exited.
func (s *BaseCQLParserListener) ExitIndexColumnSpec(ctx *IndexColumnSpecContext) {}

// EnterIndexKeysSpec is called when production indexKeysSpec is entered.
func (s *BaseCQLParserListener) EnterIndexKeysSpec(ctx *IndexKeysSpecContext) {}

// ExitIndexKeysSpec is called when production indexKeysSpec is exited.
func (s *BaseCQLParserListener) ExitIndexKeysSpec(ctx *IndexKeysSpecContext) {}

// EnterIndexEntriesSSpec is called when production indexEntriesSSpec is entered.
func (s *BaseCQLParserListener) EnterIndexEntriesSSpec(ctx *IndexEntriesSSpecContext) {}

// ExitIndexEntriesSSpec is called when production indexEntriesSSpec is exited.
func (s *BaseCQLParserListener) ExitIndexEntriesSSpec(ctx *IndexEntriesSSpecContext) {}

// EnterIndexFullSpec is called when production indexFullSpec is entered.
func (s *BaseCQLParserListener) EnterIndexFullSpec(ctx *IndexFullSpecContext) {}

// ExitIndexFullSpec is called when production indexFullSpec is exited.
func (s *BaseCQLParserListener) ExitIndexFullSpec(ctx *IndexFullSpecContext) {}

// EnterDelete_ is called when production delete_ is entered.
func (s *BaseCQLParserListener) EnterDelete_(ctx *Delete_Context) {}

// ExitDelete_ is called when production delete_ is exited.
func (s *BaseCQLParserListener) ExitDelete_(ctx *Delete_Context) {}

// EnterDeleteColumnList is called when production deleteColumnList is entered.
func (s *BaseCQLParserListener) EnterDeleteColumnList(ctx *DeleteColumnListContext) {}

// ExitDeleteColumnList is called when production deleteColumnList is exited.
func (s *BaseCQLParserListener) ExitDeleteColumnList(ctx *DeleteColumnListContext) {}

// EnterDeleteColumnItem is called when production deleteColumnItem is entered.
func (s *BaseCQLParserListener) EnterDeleteColumnItem(ctx *DeleteColumnItemContext) {}

// ExitDeleteColumnItem is called when production deleteColumnItem is exited.
func (s *BaseCQLParserListener) ExitDeleteColumnItem(ctx *DeleteColumnItemContext) {}

// EnterUpdate is called when production update is entered.
func (s *BaseCQLParserListener) EnterUpdate(ctx *UpdateContext) {}

// ExitUpdate is called when production update is exited.
func (s *BaseCQLParserListener) ExitUpdate(ctx *UpdateContext) {}

// EnterIfSpec is called when production ifSpec is entered.
func (s *BaseCQLParserListener) EnterIfSpec(ctx *IfSpecContext) {}

// ExitIfSpec is called when production ifSpec is exited.
func (s *BaseCQLParserListener) ExitIfSpec(ctx *IfSpecContext) {}

// EnterIfConditionList is called when production ifConditionList is entered.
func (s *BaseCQLParserListener) EnterIfConditionList(ctx *IfConditionListContext) {}

// ExitIfConditionList is called when production ifConditionList is exited.
func (s *BaseCQLParserListener) ExitIfConditionList(ctx *IfConditionListContext) {}

// EnterIfCondition is called when production ifCondition is entered.
func (s *BaseCQLParserListener) EnterIfCondition(ctx *IfConditionContext) {}

// ExitIfCondition is called when production ifCondition is exited.
func (s *BaseCQLParserListener) ExitIfCondition(ctx *IfConditionContext) {}

// EnterAssignments is called when production assignments is entered.
func (s *BaseCQLParserListener) EnterAssignments(ctx *AssignmentsContext) {}

// ExitAssignments is called when production assignments is exited.
func (s *BaseCQLParserListener) ExitAssignments(ctx *AssignmentsContext) {}

// EnterAssignmentElement is called when production assignmentElement is entered.
func (s *BaseCQLParserListener) EnterAssignmentElement(ctx *AssignmentElementContext) {}

// ExitAssignmentElement is called when production assignmentElement is exited.
func (s *BaseCQLParserListener) ExitAssignmentElement(ctx *AssignmentElementContext) {}

// EnterAssignmentSet is called when production assignmentSet is entered.
func (s *BaseCQLParserListener) EnterAssignmentSet(ctx *AssignmentSetContext) {}

// ExitAssignmentSet is called when production assignmentSet is exited.
func (s *BaseCQLParserListener) ExitAssignmentSet(ctx *AssignmentSetContext) {}

// EnterAssignmentMap is called when production assignmentMap is entered.
func (s *BaseCQLParserListener) EnterAssignmentMap(ctx *AssignmentMapContext) {}

// ExitAssignmentMap is called when production assignmentMap is exited.
func (s *BaseCQLParserListener) ExitAssignmentMap(ctx *AssignmentMapContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseCQLParserListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseCQLParserListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterAssignmentTuple is called when production assignmentTuple is entered.
func (s *BaseCQLParserListener) EnterAssignmentTuple(ctx *AssignmentTupleContext) {}

// ExitAssignmentTuple is called when production assignmentTuple is exited.
func (s *BaseCQLParserListener) ExitAssignmentTuple(ctx *AssignmentTupleContext) {}

// EnterInsert is called when production insert is entered.
func (s *BaseCQLParserListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production insert is exited.
func (s *BaseCQLParserListener) ExitInsert(ctx *InsertContext) {}

// EnterUsingTtlTimestamp is called when production usingTtlTimestamp is entered.
func (s *BaseCQLParserListener) EnterUsingTtlTimestamp(ctx *UsingTtlTimestampContext) {}

// ExitUsingTtlTimestamp is called when production usingTtlTimestamp is exited.
func (s *BaseCQLParserListener) ExitUsingTtlTimestamp(ctx *UsingTtlTimestampContext) {}

// EnterTimestamp is called when production timestamp is entered.
func (s *BaseCQLParserListener) EnterTimestamp(ctx *TimestampContext) {}

// ExitTimestamp is called when production timestamp is exited.
func (s *BaseCQLParserListener) ExitTimestamp(ctx *TimestampContext) {}

// EnterTtl is called when production ttl is entered.
func (s *BaseCQLParserListener) EnterTtl(ctx *TtlContext) {}

// ExitTtl is called when production ttl is exited.
func (s *BaseCQLParserListener) ExitTtl(ctx *TtlContext) {}

// EnterUsingTimestampSpec is called when production usingTimestampSpec is entered.
func (s *BaseCQLParserListener) EnterUsingTimestampSpec(ctx *UsingTimestampSpecContext) {}

// ExitUsingTimestampSpec is called when production usingTimestampSpec is exited.
func (s *BaseCQLParserListener) ExitUsingTimestampSpec(ctx *UsingTimestampSpecContext) {}

// EnterIfNotExist is called when production ifNotExist is entered.
func (s *BaseCQLParserListener) EnterIfNotExist(ctx *IfNotExistContext) {}

// ExitIfNotExist is called when production ifNotExist is exited.
func (s *BaseCQLParserListener) ExitIfNotExist(ctx *IfNotExistContext) {}

// EnterIfExist is called when production ifExist is entered.
func (s *BaseCQLParserListener) EnterIfExist(ctx *IfExistContext) {}

// ExitIfExist is called when production ifExist is exited.
func (s *BaseCQLParserListener) ExitIfExist(ctx *IfExistContext) {}

// EnterInsertValuesSpec is called when production insertValuesSpec is entered.
func (s *BaseCQLParserListener) EnterInsertValuesSpec(ctx *InsertValuesSpecContext) {}

// ExitInsertValuesSpec is called when production insertValuesSpec is exited.
func (s *BaseCQLParserListener) ExitInsertValuesSpec(ctx *InsertValuesSpecContext) {}

// EnterInsertColumnSpec is called when production insertColumnSpec is entered.
func (s *BaseCQLParserListener) EnterInsertColumnSpec(ctx *InsertColumnSpecContext) {}

// ExitInsertColumnSpec is called when production insertColumnSpec is exited.
func (s *BaseCQLParserListener) ExitInsertColumnSpec(ctx *InsertColumnSpecContext) {}

// EnterColumnList is called when production columnList is entered.
func (s *BaseCQLParserListener) EnterColumnList(ctx *ColumnListContext) {}

// ExitColumnList is called when production columnList is exited.
func (s *BaseCQLParserListener) ExitColumnList(ctx *ColumnListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseCQLParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseCQLParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCQLParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCQLParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSelect_ is called when production select_ is entered.
func (s *BaseCQLParserListener) EnterSelect_(ctx *Select_Context) {}

// ExitSelect_ is called when production select_ is exited.
func (s *BaseCQLParserListener) ExitSelect_(ctx *Select_Context) {}

// EnterAllowFilteringSpec is called when production allowFilteringSpec is entered.
func (s *BaseCQLParserListener) EnterAllowFilteringSpec(ctx *AllowFilteringSpecContext) {}

// ExitAllowFilteringSpec is called when production allowFilteringSpec is exited.
func (s *BaseCQLParserListener) ExitAllowFilteringSpec(ctx *AllowFilteringSpecContext) {}

// EnterLimitSpec is called when production limitSpec is entered.
func (s *BaseCQLParserListener) EnterLimitSpec(ctx *LimitSpecContext) {}

// ExitLimitSpec is called when production limitSpec is exited.
func (s *BaseCQLParserListener) ExitLimitSpec(ctx *LimitSpecContext) {}

// EnterFromSpec is called when production fromSpec is entered.
func (s *BaseCQLParserListener) EnterFromSpec(ctx *FromSpecContext) {}

// ExitFromSpec is called when production fromSpec is exited.
func (s *BaseCQLParserListener) ExitFromSpec(ctx *FromSpecContext) {}

// EnterFromSpecElement is called when production fromSpecElement is entered.
func (s *BaseCQLParserListener) EnterFromSpecElement(ctx *FromSpecElementContext) {}

// ExitFromSpecElement is called when production fromSpecElement is exited.
func (s *BaseCQLParserListener) ExitFromSpecElement(ctx *FromSpecElementContext) {}

// EnterOrderSpec is called when production orderSpec is entered.
func (s *BaseCQLParserListener) EnterOrderSpec(ctx *OrderSpecContext) {}

// ExitOrderSpec is called when production orderSpec is exited.
func (s *BaseCQLParserListener) ExitOrderSpec(ctx *OrderSpecContext) {}

// EnterOrderSpecElement is called when production orderSpecElement is entered.
func (s *BaseCQLParserListener) EnterOrderSpecElement(ctx *OrderSpecElementContext) {}

// ExitOrderSpecElement is called when production orderSpecElement is exited.
func (s *BaseCQLParserListener) ExitOrderSpecElement(ctx *OrderSpecElementContext) {}

// EnterWhereSpec is called when production whereSpec is entered.
func (s *BaseCQLParserListener) EnterWhereSpec(ctx *WhereSpecContext) {}

// ExitWhereSpec is called when production whereSpec is exited.
func (s *BaseCQLParserListener) ExitWhereSpec(ctx *WhereSpecContext) {}

// EnterDistinctSpec is called when production distinctSpec is entered.
func (s *BaseCQLParserListener) EnterDistinctSpec(ctx *DistinctSpecContext) {}

// ExitDistinctSpec is called when production distinctSpec is exited.
func (s *BaseCQLParserListener) ExitDistinctSpec(ctx *DistinctSpecContext) {}

// EnterSelectElements is called when production selectElements is entered.
func (s *BaseCQLParserListener) EnterSelectElements(ctx *SelectElementsContext) {}

// ExitSelectElements is called when production selectElements is exited.
func (s *BaseCQLParserListener) ExitSelectElements(ctx *SelectElementsContext) {}

// EnterSelectElement is called when production selectElement is entered.
func (s *BaseCQLParserListener) EnterSelectElement(ctx *SelectElementContext) {}

// ExitSelectElement is called when production selectElement is exited.
func (s *BaseCQLParserListener) ExitSelectElement(ctx *SelectElementContext) {}

// EnterRelationElements is called when production relationElements is entered.
func (s *BaseCQLParserListener) EnterRelationElements(ctx *RelationElementsContext) {}

// ExitRelationElements is called when production relationElements is exited.
func (s *BaseCQLParserListener) ExitRelationElements(ctx *RelationElementsContext) {}

// EnterRelationElement is called when production relationElement is entered.
func (s *BaseCQLParserListener) EnterRelationElement(ctx *RelationElementContext) {}

// ExitRelationElement is called when production relationElement is exited.
func (s *BaseCQLParserListener) ExitRelationElement(ctx *RelationElementContext) {}

// EnterRelalationContains is called when production relalationContains is entered.
func (s *BaseCQLParserListener) EnterRelalationContains(ctx *RelalationContainsContext) {}

// ExitRelalationContains is called when production relalationContains is exited.
func (s *BaseCQLParserListener) ExitRelalationContains(ctx *RelalationContainsContext) {}

// EnterRelalationContainsKey is called when production relalationContainsKey is entered.
func (s *BaseCQLParserListener) EnterRelalationContainsKey(ctx *RelalationContainsKeyContext) {}

// ExitRelalationContainsKey is called when production relalationContainsKey is exited.
func (s *BaseCQLParserListener) ExitRelalationContainsKey(ctx *RelalationContainsKeyContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseCQLParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseCQLParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BaseCQLParserListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BaseCQLParserListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCQLParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCQLParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseCQLParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseCQLParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseCQLParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseCQLParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseCQLParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseCQLParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseCQLParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseCQLParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterHexadecimalLiteral is called when production hexadecimalLiteral is entered.
func (s *BaseCQLParserListener) EnterHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// ExitHexadecimalLiteral is called when production hexadecimalLiteral is exited.
func (s *BaseCQLParserListener) ExitHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// EnterKeyspace is called when production keyspace is entered.
func (s *BaseCQLParserListener) EnterKeyspace(ctx *KeyspaceContext) {}

// ExitKeyspace is called when production keyspace is exited.
func (s *BaseCQLParserListener) ExitKeyspace(ctx *KeyspaceContext) {}

// EnterTable is called when production table is entered.
func (s *BaseCQLParserListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseCQLParserListener) ExitTable(ctx *TableContext) {}

// EnterColumn is called when production column is entered.
func (s *BaseCQLParserListener) EnterColumn(ctx *ColumnContext) {}

// ExitColumn is called when production column is exited.
func (s *BaseCQLParserListener) ExitColumn(ctx *ColumnContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseCQLParserListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseCQLParserListener) ExitDataType(ctx *DataTypeContext) {}

// EnterDataTypeName is called when production dataTypeName is entered.
func (s *BaseCQLParserListener) EnterDataTypeName(ctx *DataTypeNameContext) {}

// ExitDataTypeName is called when production dataTypeName is exited.
func (s *BaseCQLParserListener) ExitDataTypeName(ctx *DataTypeNameContext) {}

// EnterDataTypeDefinition is called when production dataTypeDefinition is entered.
func (s *BaseCQLParserListener) EnterDataTypeDefinition(ctx *DataTypeDefinitionContext) {}

// ExitDataTypeDefinition is called when production dataTypeDefinition is exited.
func (s *BaseCQLParserListener) ExitDataTypeDefinition(ctx *DataTypeDefinitionContext) {}

// EnterOrderDirection is called when production orderDirection is entered.
func (s *BaseCQLParserListener) EnterOrderDirection(ctx *OrderDirectionContext) {}

// ExitOrderDirection is called when production orderDirection is exited.
func (s *BaseCQLParserListener) ExitOrderDirection(ctx *OrderDirectionContext) {}

// EnterRole is called when production role is entered.
func (s *BaseCQLParserListener) EnterRole(ctx *RoleContext) {}

// ExitRole is called when production role is exited.
func (s *BaseCQLParserListener) ExitRole(ctx *RoleContext) {}

// EnterTrigger is called when production trigger is entered.
func (s *BaseCQLParserListener) EnterTrigger(ctx *TriggerContext) {}

// ExitTrigger is called when production trigger is exited.
func (s *BaseCQLParserListener) ExitTrigger(ctx *TriggerContext) {}

// EnterTriggerClass is called when production triggerClass is entered.
func (s *BaseCQLParserListener) EnterTriggerClass(ctx *TriggerClassContext) {}

// ExitTriggerClass is called when production triggerClass is exited.
func (s *BaseCQLParserListener) ExitTriggerClass(ctx *TriggerClassContext) {}

// EnterMaterializedView is called when production materializedView is entered.
func (s *BaseCQLParserListener) EnterMaterializedView(ctx *MaterializedViewContext) {}

// ExitMaterializedView is called when production materializedView is exited.
func (s *BaseCQLParserListener) ExitMaterializedView(ctx *MaterializedViewContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseCQLParserListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseCQLParserListener) ExitType_(ctx *Type_Context) {}

// EnterAggregate is called when production aggregate is entered.
func (s *BaseCQLParserListener) EnterAggregate(ctx *AggregateContext) {}

// ExitAggregate is called when production aggregate is exited.
func (s *BaseCQLParserListener) ExitAggregate(ctx *AggregateContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BaseCQLParserListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BaseCQLParserListener) ExitFunction_(ctx *Function_Context) {}

// EnterLanguage is called when production language is entered.
func (s *BaseCQLParserListener) EnterLanguage(ctx *LanguageContext) {}

// ExitLanguage is called when production language is exited.
func (s *BaseCQLParserListener) ExitLanguage(ctx *LanguageContext) {}

// EnterUser is called when production user is entered.
func (s *BaseCQLParserListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseCQLParserListener) ExitUser(ctx *UserContext) {}

// EnterPassword is called when production password is entered.
func (s *BaseCQLParserListener) EnterPassword(ctx *PasswordContext) {}

// ExitPassword is called when production password is exited.
func (s *BaseCQLParserListener) ExitPassword(ctx *PasswordContext) {}

// EnterHashKey is called when production hashKey is entered.
func (s *BaseCQLParserListener) EnterHashKey(ctx *HashKeyContext) {}

// ExitHashKey is called when production hashKey is exited.
func (s *BaseCQLParserListener) ExitHashKey(ctx *HashKeyContext) {}

// EnterParam is called when production param is entered.
func (s *BaseCQLParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseCQLParserListener) ExitParam(ctx *ParamContext) {}

// EnterParamName is called when production paramName is entered.
func (s *BaseCQLParserListener) EnterParamName(ctx *ParamNameContext) {}

// ExitParamName is called when production paramName is exited.
func (s *BaseCQLParserListener) ExitParamName(ctx *ParamNameContext) {}

// EnterKwAdd is called when production kwAdd is entered.
func (s *BaseCQLParserListener) EnterKwAdd(ctx *KwAddContext) {}

// ExitKwAdd is called when production kwAdd is exited.
func (s *BaseCQLParserListener) ExitKwAdd(ctx *KwAddContext) {}

// EnterKwAggregate is called when production kwAggregate is entered.
func (s *BaseCQLParserListener) EnterKwAggregate(ctx *KwAggregateContext) {}

// ExitKwAggregate is called when production kwAggregate is exited.
func (s *BaseCQLParserListener) ExitKwAggregate(ctx *KwAggregateContext) {}

// EnterKwAll is called when production kwAll is entered.
func (s *BaseCQLParserListener) EnterKwAll(ctx *KwAllContext) {}

// ExitKwAll is called when production kwAll is exited.
func (s *BaseCQLParserListener) ExitKwAll(ctx *KwAllContext) {}

// EnterKwAllPermissions is called when production kwAllPermissions is entered.
func (s *BaseCQLParserListener) EnterKwAllPermissions(ctx *KwAllPermissionsContext) {}

// ExitKwAllPermissions is called when production kwAllPermissions is exited.
func (s *BaseCQLParserListener) ExitKwAllPermissions(ctx *KwAllPermissionsContext) {}

// EnterKwAllow is called when production kwAllow is entered.
func (s *BaseCQLParserListener) EnterKwAllow(ctx *KwAllowContext) {}

// ExitKwAllow is called when production kwAllow is exited.
func (s *BaseCQLParserListener) ExitKwAllow(ctx *KwAllowContext) {}

// EnterKwAlter is called when production kwAlter is entered.
func (s *BaseCQLParserListener) EnterKwAlter(ctx *KwAlterContext) {}

// ExitKwAlter is called when production kwAlter is exited.
func (s *BaseCQLParserListener) ExitKwAlter(ctx *KwAlterContext) {}

// EnterKwAnd is called when production kwAnd is entered.
func (s *BaseCQLParserListener) EnterKwAnd(ctx *KwAndContext) {}

// ExitKwAnd is called when production kwAnd is exited.
func (s *BaseCQLParserListener) ExitKwAnd(ctx *KwAndContext) {}

// EnterKwApply is called when production kwApply is entered.
func (s *BaseCQLParserListener) EnterKwApply(ctx *KwApplyContext) {}

// ExitKwApply is called when production kwApply is exited.
func (s *BaseCQLParserListener) ExitKwApply(ctx *KwApplyContext) {}

// EnterKwAs is called when production kwAs is entered.
func (s *BaseCQLParserListener) EnterKwAs(ctx *KwAsContext) {}

// ExitKwAs is called when production kwAs is exited.
func (s *BaseCQLParserListener) ExitKwAs(ctx *KwAsContext) {}

// EnterKwAsc is called when production kwAsc is entered.
func (s *BaseCQLParserListener) EnterKwAsc(ctx *KwAscContext) {}

// ExitKwAsc is called when production kwAsc is exited.
func (s *BaseCQLParserListener) ExitKwAsc(ctx *KwAscContext) {}

// EnterKwAuthorize is called when production kwAuthorize is entered.
func (s *BaseCQLParserListener) EnterKwAuthorize(ctx *KwAuthorizeContext) {}

// ExitKwAuthorize is called when production kwAuthorize is exited.
func (s *BaseCQLParserListener) ExitKwAuthorize(ctx *KwAuthorizeContext) {}

// EnterKwBatch is called when production kwBatch is entered.
func (s *BaseCQLParserListener) EnterKwBatch(ctx *KwBatchContext) {}

// ExitKwBatch is called when production kwBatch is exited.
func (s *BaseCQLParserListener) ExitKwBatch(ctx *KwBatchContext) {}

// EnterKwBegin is called when production kwBegin is entered.
func (s *BaseCQLParserListener) EnterKwBegin(ctx *KwBeginContext) {}

// ExitKwBegin is called when production kwBegin is exited.
func (s *BaseCQLParserListener) ExitKwBegin(ctx *KwBeginContext) {}

// EnterKwBy is called when production kwBy is entered.
func (s *BaseCQLParserListener) EnterKwBy(ctx *KwByContext) {}

// ExitKwBy is called when production kwBy is exited.
func (s *BaseCQLParserListener) ExitKwBy(ctx *KwByContext) {}

// EnterKwCalled is called when production kwCalled is entered.
func (s *BaseCQLParserListener) EnterKwCalled(ctx *KwCalledContext) {}

// ExitKwCalled is called when production kwCalled is exited.
func (s *BaseCQLParserListener) ExitKwCalled(ctx *KwCalledContext) {}

// EnterKwClustering is called when production kwClustering is entered.
func (s *BaseCQLParserListener) EnterKwClustering(ctx *KwClusteringContext) {}

// ExitKwClustering is called when production kwClustering is exited.
func (s *BaseCQLParserListener) ExitKwClustering(ctx *KwClusteringContext) {}

// EnterKwCompact is called when production kwCompact is entered.
func (s *BaseCQLParserListener) EnterKwCompact(ctx *KwCompactContext) {}

// ExitKwCompact is called when production kwCompact is exited.
func (s *BaseCQLParserListener) ExitKwCompact(ctx *KwCompactContext) {}

// EnterKwContains is called when production kwContains is entered.
func (s *BaseCQLParserListener) EnterKwContains(ctx *KwContainsContext) {}

// ExitKwContains is called when production kwContains is exited.
func (s *BaseCQLParserListener) ExitKwContains(ctx *KwContainsContext) {}

// EnterKwCreate is called when production kwCreate is entered.
func (s *BaseCQLParserListener) EnterKwCreate(ctx *KwCreateContext) {}

// ExitKwCreate is called when production kwCreate is exited.
func (s *BaseCQLParserListener) ExitKwCreate(ctx *KwCreateContext) {}

// EnterKwDelete is called when production kwDelete is entered.
func (s *BaseCQLParserListener) EnterKwDelete(ctx *KwDeleteContext) {}

// ExitKwDelete is called when production kwDelete is exited.
func (s *BaseCQLParserListener) ExitKwDelete(ctx *KwDeleteContext) {}

// EnterKwDesc is called when production kwDesc is entered.
func (s *BaseCQLParserListener) EnterKwDesc(ctx *KwDescContext) {}

// ExitKwDesc is called when production kwDesc is exited.
func (s *BaseCQLParserListener) ExitKwDesc(ctx *KwDescContext) {}

// EnterKwDescibe is called when production kwDescibe is entered.
func (s *BaseCQLParserListener) EnterKwDescibe(ctx *KwDescibeContext) {}

// ExitKwDescibe is called when production kwDescibe is exited.
func (s *BaseCQLParserListener) ExitKwDescibe(ctx *KwDescibeContext) {}

// EnterKwDistinct is called when production kwDistinct is entered.
func (s *BaseCQLParserListener) EnterKwDistinct(ctx *KwDistinctContext) {}

// ExitKwDistinct is called when production kwDistinct is exited.
func (s *BaseCQLParserListener) ExitKwDistinct(ctx *KwDistinctContext) {}

// EnterKwDrop is called when production kwDrop is entered.
func (s *BaseCQLParserListener) EnterKwDrop(ctx *KwDropContext) {}

// ExitKwDrop is called when production kwDrop is exited.
func (s *BaseCQLParserListener) ExitKwDrop(ctx *KwDropContext) {}

// EnterKwDurableWrites is called when production kwDurableWrites is entered.
func (s *BaseCQLParserListener) EnterKwDurableWrites(ctx *KwDurableWritesContext) {}

// ExitKwDurableWrites is called when production kwDurableWrites is exited.
func (s *BaseCQLParserListener) ExitKwDurableWrites(ctx *KwDurableWritesContext) {}

// EnterKwEntries is called when production kwEntries is entered.
func (s *BaseCQLParserListener) EnterKwEntries(ctx *KwEntriesContext) {}

// ExitKwEntries is called when production kwEntries is exited.
func (s *BaseCQLParserListener) ExitKwEntries(ctx *KwEntriesContext) {}

// EnterKwExecute is called when production kwExecute is entered.
func (s *BaseCQLParserListener) EnterKwExecute(ctx *KwExecuteContext) {}

// ExitKwExecute is called when production kwExecute is exited.
func (s *BaseCQLParserListener) ExitKwExecute(ctx *KwExecuteContext) {}

// EnterKwExists is called when production kwExists is entered.
func (s *BaseCQLParserListener) EnterKwExists(ctx *KwExistsContext) {}

// ExitKwExists is called when production kwExists is exited.
func (s *BaseCQLParserListener) ExitKwExists(ctx *KwExistsContext) {}

// EnterKwFiltering is called when production kwFiltering is entered.
func (s *BaseCQLParserListener) EnterKwFiltering(ctx *KwFilteringContext) {}

// ExitKwFiltering is called when production kwFiltering is exited.
func (s *BaseCQLParserListener) ExitKwFiltering(ctx *KwFilteringContext) {}

// EnterKwFinalfunc is called when production kwFinalfunc is entered.
func (s *BaseCQLParserListener) EnterKwFinalfunc(ctx *KwFinalfuncContext) {}

// ExitKwFinalfunc is called when production kwFinalfunc is exited.
func (s *BaseCQLParserListener) ExitKwFinalfunc(ctx *KwFinalfuncContext) {}

// EnterKwFrom is called when production kwFrom is entered.
func (s *BaseCQLParserListener) EnterKwFrom(ctx *KwFromContext) {}

// ExitKwFrom is called when production kwFrom is exited.
func (s *BaseCQLParserListener) ExitKwFrom(ctx *KwFromContext) {}

// EnterKwFull is called when production kwFull is entered.
func (s *BaseCQLParserListener) EnterKwFull(ctx *KwFullContext) {}

// ExitKwFull is called when production kwFull is exited.
func (s *BaseCQLParserListener) ExitKwFull(ctx *KwFullContext) {}

// EnterKwFunction is called when production kwFunction is entered.
func (s *BaseCQLParserListener) EnterKwFunction(ctx *KwFunctionContext) {}

// ExitKwFunction is called when production kwFunction is exited.
func (s *BaseCQLParserListener) ExitKwFunction(ctx *KwFunctionContext) {}

// EnterKwFunctions is called when production kwFunctions is entered.
func (s *BaseCQLParserListener) EnterKwFunctions(ctx *KwFunctionsContext) {}

// ExitKwFunctions is called when production kwFunctions is exited.
func (s *BaseCQLParserListener) ExitKwFunctions(ctx *KwFunctionsContext) {}

// EnterKwGrant is called when production kwGrant is entered.
func (s *BaseCQLParserListener) EnterKwGrant(ctx *KwGrantContext) {}

// ExitKwGrant is called when production kwGrant is exited.
func (s *BaseCQLParserListener) ExitKwGrant(ctx *KwGrantContext) {}

// EnterKwIf is called when production kwIf is entered.
func (s *BaseCQLParserListener) EnterKwIf(ctx *KwIfContext) {}

// ExitKwIf is called when production kwIf is exited.
func (s *BaseCQLParserListener) ExitKwIf(ctx *KwIfContext) {}

// EnterKwIn is called when production kwIn is entered.
func (s *BaseCQLParserListener) EnterKwIn(ctx *KwInContext) {}

// ExitKwIn is called when production kwIn is exited.
func (s *BaseCQLParserListener) ExitKwIn(ctx *KwInContext) {}

// EnterKwIndex is called when production kwIndex is entered.
func (s *BaseCQLParserListener) EnterKwIndex(ctx *KwIndexContext) {}

// ExitKwIndex is called when production kwIndex is exited.
func (s *BaseCQLParserListener) ExitKwIndex(ctx *KwIndexContext) {}

// EnterKwInitcond is called when production kwInitcond is entered.
func (s *BaseCQLParserListener) EnterKwInitcond(ctx *KwInitcondContext) {}

// ExitKwInitcond is called when production kwInitcond is exited.
func (s *BaseCQLParserListener) ExitKwInitcond(ctx *KwInitcondContext) {}

// EnterKwInput is called when production kwInput is entered.
func (s *BaseCQLParserListener) EnterKwInput(ctx *KwInputContext) {}

// ExitKwInput is called when production kwInput is exited.
func (s *BaseCQLParserListener) ExitKwInput(ctx *KwInputContext) {}

// EnterKwInsert is called when production kwInsert is entered.
func (s *BaseCQLParserListener) EnterKwInsert(ctx *KwInsertContext) {}

// ExitKwInsert is called when production kwInsert is exited.
func (s *BaseCQLParserListener) ExitKwInsert(ctx *KwInsertContext) {}

// EnterKwInto is called when production kwInto is entered.
func (s *BaseCQLParserListener) EnterKwInto(ctx *KwIntoContext) {}

// ExitKwInto is called when production kwInto is exited.
func (s *BaseCQLParserListener) ExitKwInto(ctx *KwIntoContext) {}

// EnterKwIs is called when production kwIs is entered.
func (s *BaseCQLParserListener) EnterKwIs(ctx *KwIsContext) {}

// ExitKwIs is called when production kwIs is exited.
func (s *BaseCQLParserListener) ExitKwIs(ctx *KwIsContext) {}

// EnterKwJson is called when production kwJson is entered.
func (s *BaseCQLParserListener) EnterKwJson(ctx *KwJsonContext) {}

// ExitKwJson is called when production kwJson is exited.
func (s *BaseCQLParserListener) ExitKwJson(ctx *KwJsonContext) {}

// EnterKwKey is called when production kwKey is entered.
func (s *BaseCQLParserListener) EnterKwKey(ctx *KwKeyContext) {}

// ExitKwKey is called when production kwKey is exited.
func (s *BaseCQLParserListener) ExitKwKey(ctx *KwKeyContext) {}

// EnterKwKeys is called when production kwKeys is entered.
func (s *BaseCQLParserListener) EnterKwKeys(ctx *KwKeysContext) {}

// ExitKwKeys is called when production kwKeys is exited.
func (s *BaseCQLParserListener) ExitKwKeys(ctx *KwKeysContext) {}

// EnterKwKeyspace is called when production kwKeyspace is entered.
func (s *BaseCQLParserListener) EnterKwKeyspace(ctx *KwKeyspaceContext) {}

// ExitKwKeyspace is called when production kwKeyspace is exited.
func (s *BaseCQLParserListener) ExitKwKeyspace(ctx *KwKeyspaceContext) {}

// EnterKwKeyspaces is called when production kwKeyspaces is entered.
func (s *BaseCQLParserListener) EnterKwKeyspaces(ctx *KwKeyspacesContext) {}

// ExitKwKeyspaces is called when production kwKeyspaces is exited.
func (s *BaseCQLParserListener) ExitKwKeyspaces(ctx *KwKeyspacesContext) {}

// EnterKwLanguage is called when production kwLanguage is entered.
func (s *BaseCQLParserListener) EnterKwLanguage(ctx *KwLanguageContext) {}

// ExitKwLanguage is called when production kwLanguage is exited.
func (s *BaseCQLParserListener) ExitKwLanguage(ctx *KwLanguageContext) {}

// EnterKwLimit is called when production kwLimit is entered.
func (s *BaseCQLParserListener) EnterKwLimit(ctx *KwLimitContext) {}

// ExitKwLimit is called when production kwLimit is exited.
func (s *BaseCQLParserListener) ExitKwLimit(ctx *KwLimitContext) {}

// EnterKwList is called when production kwList is entered.
func (s *BaseCQLParserListener) EnterKwList(ctx *KwListContext) {}

// ExitKwList is called when production kwList is exited.
func (s *BaseCQLParserListener) ExitKwList(ctx *KwListContext) {}

// EnterKwLogged is called when production kwLogged is entered.
func (s *BaseCQLParserListener) EnterKwLogged(ctx *KwLoggedContext) {}

// ExitKwLogged is called when production kwLogged is exited.
func (s *BaseCQLParserListener) ExitKwLogged(ctx *KwLoggedContext) {}

// EnterKwLogin is called when production kwLogin is entered.
func (s *BaseCQLParserListener) EnterKwLogin(ctx *KwLoginContext) {}

// ExitKwLogin is called when production kwLogin is exited.
func (s *BaseCQLParserListener) ExitKwLogin(ctx *KwLoginContext) {}

// EnterKwMaterialized is called when production kwMaterialized is entered.
func (s *BaseCQLParserListener) EnterKwMaterialized(ctx *KwMaterializedContext) {}

// ExitKwMaterialized is called when production kwMaterialized is exited.
func (s *BaseCQLParserListener) ExitKwMaterialized(ctx *KwMaterializedContext) {}

// EnterKwModify is called when production kwModify is entered.
func (s *BaseCQLParserListener) EnterKwModify(ctx *KwModifyContext) {}

// ExitKwModify is called when production kwModify is exited.
func (s *BaseCQLParserListener) ExitKwModify(ctx *KwModifyContext) {}

// EnterKwNosuperuser is called when production kwNosuperuser is entered.
func (s *BaseCQLParserListener) EnterKwNosuperuser(ctx *KwNosuperuserContext) {}

// ExitKwNosuperuser is called when production kwNosuperuser is exited.
func (s *BaseCQLParserListener) ExitKwNosuperuser(ctx *KwNosuperuserContext) {}

// EnterKwNorecursive is called when production kwNorecursive is entered.
func (s *BaseCQLParserListener) EnterKwNorecursive(ctx *KwNorecursiveContext) {}

// ExitKwNorecursive is called when production kwNorecursive is exited.
func (s *BaseCQLParserListener) ExitKwNorecursive(ctx *KwNorecursiveContext) {}

// EnterKwNot is called when production kwNot is entered.
func (s *BaseCQLParserListener) EnterKwNot(ctx *KwNotContext) {}

// ExitKwNot is called when production kwNot is exited.
func (s *BaseCQLParserListener) ExitKwNot(ctx *KwNotContext) {}

// EnterKwNull is called when production kwNull is entered.
func (s *BaseCQLParserListener) EnterKwNull(ctx *KwNullContext) {}

// ExitKwNull is called when production kwNull is exited.
func (s *BaseCQLParserListener) ExitKwNull(ctx *KwNullContext) {}

// EnterKwOf is called when production kwOf is entered.
func (s *BaseCQLParserListener) EnterKwOf(ctx *KwOfContext) {}

// ExitKwOf is called when production kwOf is exited.
func (s *BaseCQLParserListener) ExitKwOf(ctx *KwOfContext) {}

// EnterKwOn is called when production kwOn is entered.
func (s *BaseCQLParserListener) EnterKwOn(ctx *KwOnContext) {}

// ExitKwOn is called when production kwOn is exited.
func (s *BaseCQLParserListener) ExitKwOn(ctx *KwOnContext) {}

// EnterKwOptions is called when production kwOptions is entered.
func (s *BaseCQLParserListener) EnterKwOptions(ctx *KwOptionsContext) {}

// ExitKwOptions is called when production kwOptions is exited.
func (s *BaseCQLParserListener) ExitKwOptions(ctx *KwOptionsContext) {}

// EnterKwOr is called when production kwOr is entered.
func (s *BaseCQLParserListener) EnterKwOr(ctx *KwOrContext) {}

// ExitKwOr is called when production kwOr is exited.
func (s *BaseCQLParserListener) ExitKwOr(ctx *KwOrContext) {}

// EnterKwOrder is called when production kwOrder is entered.
func (s *BaseCQLParserListener) EnterKwOrder(ctx *KwOrderContext) {}

// ExitKwOrder is called when production kwOrder is exited.
func (s *BaseCQLParserListener) ExitKwOrder(ctx *KwOrderContext) {}

// EnterKwPassword is called when production kwPassword is entered.
func (s *BaseCQLParserListener) EnterKwPassword(ctx *KwPasswordContext) {}

// ExitKwPassword is called when production kwPassword is exited.
func (s *BaseCQLParserListener) ExitKwPassword(ctx *KwPasswordContext) {}

// EnterKwPrimary is called when production kwPrimary is entered.
func (s *BaseCQLParserListener) EnterKwPrimary(ctx *KwPrimaryContext) {}

// ExitKwPrimary is called when production kwPrimary is exited.
func (s *BaseCQLParserListener) ExitKwPrimary(ctx *KwPrimaryContext) {}

// EnterKwRename is called when production kwRename is entered.
func (s *BaseCQLParserListener) EnterKwRename(ctx *KwRenameContext) {}

// ExitKwRename is called when production kwRename is exited.
func (s *BaseCQLParserListener) ExitKwRename(ctx *KwRenameContext) {}

// EnterKwReplace is called when production kwReplace is entered.
func (s *BaseCQLParserListener) EnterKwReplace(ctx *KwReplaceContext) {}

// ExitKwReplace is called when production kwReplace is exited.
func (s *BaseCQLParserListener) ExitKwReplace(ctx *KwReplaceContext) {}

// EnterKwReplication is called when production kwReplication is entered.
func (s *BaseCQLParserListener) EnterKwReplication(ctx *KwReplicationContext) {}

// ExitKwReplication is called when production kwReplication is exited.
func (s *BaseCQLParserListener) ExitKwReplication(ctx *KwReplicationContext) {}

// EnterKwReturns is called when production kwReturns is entered.
func (s *BaseCQLParserListener) EnterKwReturns(ctx *KwReturnsContext) {}

// ExitKwReturns is called when production kwReturns is exited.
func (s *BaseCQLParserListener) ExitKwReturns(ctx *KwReturnsContext) {}

// EnterKwRole is called when production kwRole is entered.
func (s *BaseCQLParserListener) EnterKwRole(ctx *KwRoleContext) {}

// ExitKwRole is called when production kwRole is exited.
func (s *BaseCQLParserListener) ExitKwRole(ctx *KwRoleContext) {}

// EnterKwRoles is called when production kwRoles is entered.
func (s *BaseCQLParserListener) EnterKwRoles(ctx *KwRolesContext) {}

// ExitKwRoles is called when production kwRoles is exited.
func (s *BaseCQLParserListener) ExitKwRoles(ctx *KwRolesContext) {}

// EnterKwSelect is called when production kwSelect is entered.
func (s *BaseCQLParserListener) EnterKwSelect(ctx *KwSelectContext) {}

// ExitKwSelect is called when production kwSelect is exited.
func (s *BaseCQLParserListener) ExitKwSelect(ctx *KwSelectContext) {}

// EnterKwSet is called when production kwSet is entered.
func (s *BaseCQLParserListener) EnterKwSet(ctx *KwSetContext) {}

// ExitKwSet is called when production kwSet is exited.
func (s *BaseCQLParserListener) ExitKwSet(ctx *KwSetContext) {}

// EnterKwSfunc is called when production kwSfunc is entered.
func (s *BaseCQLParserListener) EnterKwSfunc(ctx *KwSfuncContext) {}

// ExitKwSfunc is called when production kwSfunc is exited.
func (s *BaseCQLParserListener) ExitKwSfunc(ctx *KwSfuncContext) {}

// EnterKwStorage is called when production kwStorage is entered.
func (s *BaseCQLParserListener) EnterKwStorage(ctx *KwStorageContext) {}

// ExitKwStorage is called when production kwStorage is exited.
func (s *BaseCQLParserListener) ExitKwStorage(ctx *KwStorageContext) {}

// EnterKwStype is called when production kwStype is entered.
func (s *BaseCQLParserListener) EnterKwStype(ctx *KwStypeContext) {}

// ExitKwStype is called when production kwStype is exited.
func (s *BaseCQLParserListener) ExitKwStype(ctx *KwStypeContext) {}

// EnterKwSuperuser is called when production kwSuperuser is entered.
func (s *BaseCQLParserListener) EnterKwSuperuser(ctx *KwSuperuserContext) {}

// ExitKwSuperuser is called when production kwSuperuser is exited.
func (s *BaseCQLParserListener) ExitKwSuperuser(ctx *KwSuperuserContext) {}

// EnterKwTable is called when production kwTable is entered.
func (s *BaseCQLParserListener) EnterKwTable(ctx *KwTableContext) {}

// ExitKwTable is called when production kwTable is exited.
func (s *BaseCQLParserListener) ExitKwTable(ctx *KwTableContext) {}

// EnterKwTimestamp is called when production kwTimestamp is entered.
func (s *BaseCQLParserListener) EnterKwTimestamp(ctx *KwTimestampContext) {}

// ExitKwTimestamp is called when production kwTimestamp is exited.
func (s *BaseCQLParserListener) ExitKwTimestamp(ctx *KwTimestampContext) {}

// EnterKwTo is called when production kwTo is entered.
func (s *BaseCQLParserListener) EnterKwTo(ctx *KwToContext) {}

// ExitKwTo is called when production kwTo is exited.
func (s *BaseCQLParserListener) ExitKwTo(ctx *KwToContext) {}

// EnterKwTrigger is called when production kwTrigger is entered.
func (s *BaseCQLParserListener) EnterKwTrigger(ctx *KwTriggerContext) {}

// ExitKwTrigger is called when production kwTrigger is exited.
func (s *BaseCQLParserListener) ExitKwTrigger(ctx *KwTriggerContext) {}

// EnterKwTruncate is called when production kwTruncate is entered.
func (s *BaseCQLParserListener) EnterKwTruncate(ctx *KwTruncateContext) {}

// ExitKwTruncate is called when production kwTruncate is exited.
func (s *BaseCQLParserListener) ExitKwTruncate(ctx *KwTruncateContext) {}

// EnterKwTtl is called when production kwTtl is entered.
func (s *BaseCQLParserListener) EnterKwTtl(ctx *KwTtlContext) {}

// ExitKwTtl is called when production kwTtl is exited.
func (s *BaseCQLParserListener) ExitKwTtl(ctx *KwTtlContext) {}

// EnterKwType is called when production kwType is entered.
func (s *BaseCQLParserListener) EnterKwType(ctx *KwTypeContext) {}

// ExitKwType is called when production kwType is exited.
func (s *BaseCQLParserListener) ExitKwType(ctx *KwTypeContext) {}

// EnterKwUnlogged is called when production kwUnlogged is entered.
func (s *BaseCQLParserListener) EnterKwUnlogged(ctx *KwUnloggedContext) {}

// ExitKwUnlogged is called when production kwUnlogged is exited.
func (s *BaseCQLParserListener) ExitKwUnlogged(ctx *KwUnloggedContext) {}

// EnterKwUpdate is called when production kwUpdate is entered.
func (s *BaseCQLParserListener) EnterKwUpdate(ctx *KwUpdateContext) {}

// ExitKwUpdate is called when production kwUpdate is exited.
func (s *BaseCQLParserListener) ExitKwUpdate(ctx *KwUpdateContext) {}

// EnterKwUse is called when production kwUse is entered.
func (s *BaseCQLParserListener) EnterKwUse(ctx *KwUseContext) {}

// ExitKwUse is called when production kwUse is exited.
func (s *BaseCQLParserListener) ExitKwUse(ctx *KwUseContext) {}

// EnterKwUser is called when production kwUser is entered.
func (s *BaseCQLParserListener) EnterKwUser(ctx *KwUserContext) {}

// ExitKwUser is called when production kwUser is exited.
func (s *BaseCQLParserListener) ExitKwUser(ctx *KwUserContext) {}

// EnterKwUsing is called when production kwUsing is entered.
func (s *BaseCQLParserListener) EnterKwUsing(ctx *KwUsingContext) {}

// ExitKwUsing is called when production kwUsing is exited.
func (s *BaseCQLParserListener) ExitKwUsing(ctx *KwUsingContext) {}

// EnterKwValues is called when production kwValues is entered.
func (s *BaseCQLParserListener) EnterKwValues(ctx *KwValuesContext) {}

// ExitKwValues is called when production kwValues is exited.
func (s *BaseCQLParserListener) ExitKwValues(ctx *KwValuesContext) {}

// EnterKwView is called when production kwView is entered.
func (s *BaseCQLParserListener) EnterKwView(ctx *KwViewContext) {}

// ExitKwView is called when production kwView is exited.
func (s *BaseCQLParserListener) ExitKwView(ctx *KwViewContext) {}

// EnterKwWhere is called when production kwWhere is entered.
func (s *BaseCQLParserListener) EnterKwWhere(ctx *KwWhereContext) {}

// ExitKwWhere is called when production kwWhere is exited.
func (s *BaseCQLParserListener) ExitKwWhere(ctx *KwWhereContext) {}

// EnterKwWith is called when production kwWith is entered.
func (s *BaseCQLParserListener) EnterKwWith(ctx *KwWithContext) {}

// ExitKwWith is called when production kwWith is exited.
func (s *BaseCQLParserListener) ExitKwWith(ctx *KwWithContext) {}

// EnterKwRevoke is called when production kwRevoke is entered.
func (s *BaseCQLParserListener) EnterKwRevoke(ctx *KwRevokeContext) {}

// ExitKwRevoke is called when production kwRevoke is exited.
func (s *BaseCQLParserListener) ExitKwRevoke(ctx *KwRevokeContext) {}

// EnterSyntaxBracketLr is called when production syntaxBracketLr is entered.
func (s *BaseCQLParserListener) EnterSyntaxBracketLr(ctx *SyntaxBracketLrContext) {}

// ExitSyntaxBracketLr is called when production syntaxBracketLr is exited.
func (s *BaseCQLParserListener) ExitSyntaxBracketLr(ctx *SyntaxBracketLrContext) {}

// EnterSyntaxBracketRr is called when production syntaxBracketRr is entered.
func (s *BaseCQLParserListener) EnterSyntaxBracketRr(ctx *SyntaxBracketRrContext) {}

// ExitSyntaxBracketRr is called when production syntaxBracketRr is exited.
func (s *BaseCQLParserListener) ExitSyntaxBracketRr(ctx *SyntaxBracketRrContext) {}

// EnterSyntaxBracketLc is called when production syntaxBracketLc is entered.
func (s *BaseCQLParserListener) EnterSyntaxBracketLc(ctx *SyntaxBracketLcContext) {}

// ExitSyntaxBracketLc is called when production syntaxBracketLc is exited.
func (s *BaseCQLParserListener) ExitSyntaxBracketLc(ctx *SyntaxBracketLcContext) {}

// EnterSyntaxBracketRc is called when production syntaxBracketRc is entered.
func (s *BaseCQLParserListener) EnterSyntaxBracketRc(ctx *SyntaxBracketRcContext) {}

// ExitSyntaxBracketRc is called when production syntaxBracketRc is exited.
func (s *BaseCQLParserListener) ExitSyntaxBracketRc(ctx *SyntaxBracketRcContext) {}

// EnterSyntaxBracketLa is called when production syntaxBracketLa is entered.
func (s *BaseCQLParserListener) EnterSyntaxBracketLa(ctx *SyntaxBracketLaContext) {}

// ExitSyntaxBracketLa is called when production syntaxBracketLa is exited.
func (s *BaseCQLParserListener) ExitSyntaxBracketLa(ctx *SyntaxBracketLaContext) {}

// EnterSyntaxBracketRa is called when production syntaxBracketRa is entered.
func (s *BaseCQLParserListener) EnterSyntaxBracketRa(ctx *SyntaxBracketRaContext) {}

// ExitSyntaxBracketRa is called when production syntaxBracketRa is exited.
func (s *BaseCQLParserListener) ExitSyntaxBracketRa(ctx *SyntaxBracketRaContext) {}

// EnterSyntaxBracketLs is called when production syntaxBracketLs is entered.
func (s *BaseCQLParserListener) EnterSyntaxBracketLs(ctx *SyntaxBracketLsContext) {}

// ExitSyntaxBracketLs is called when production syntaxBracketLs is exited.
func (s *BaseCQLParserListener) ExitSyntaxBracketLs(ctx *SyntaxBracketLsContext) {}

// EnterSyntaxBracketRs is called when production syntaxBracketRs is entered.
func (s *BaseCQLParserListener) EnterSyntaxBracketRs(ctx *SyntaxBracketRsContext) {}

// ExitSyntaxBracketRs is called when production syntaxBracketRs is exited.
func (s *BaseCQLParserListener) ExitSyntaxBracketRs(ctx *SyntaxBracketRsContext) {}

// EnterSyntaxComma is called when production syntaxComma is entered.
func (s *BaseCQLParserListener) EnterSyntaxComma(ctx *SyntaxCommaContext) {}

// ExitSyntaxComma is called when production syntaxComma is exited.
func (s *BaseCQLParserListener) ExitSyntaxComma(ctx *SyntaxCommaContext) {}

// EnterSyntaxColon is called when production syntaxColon is entered.
func (s *BaseCQLParserListener) EnterSyntaxColon(ctx *SyntaxColonContext) {}

// ExitSyntaxColon is called when production syntaxColon is exited.
func (s *BaseCQLParserListener) ExitSyntaxColon(ctx *SyntaxColonContext) {}
