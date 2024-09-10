// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file flyteidl/admin/project.proto (package flyteidl.admin, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Labels, Sort } from "./common_pb.js";

/**
 * Empty request for GetDomain
 *
 * @generated from message flyteidl.admin.GetDomainRequest
 */
export class GetDomainRequest extends Message<GetDomainRequest> {
  constructor(data?: PartialMessage<GetDomainRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetDomainRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDomainRequest {
    return new GetDomainRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDomainRequest {
    return new GetDomainRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDomainRequest {
    return new GetDomainRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetDomainRequest | PlainMessage<GetDomainRequest> | undefined, b: GetDomainRequest | PlainMessage<GetDomainRequest> | undefined): boolean {
    return proto3.util.equals(GetDomainRequest, a, b);
  }
}

/**
 * Namespace within a project commonly used to differentiate between different service instances.
 * e.g. "production", "development", etc.
 *
 * @generated from message flyteidl.admin.Domain
 */
export class Domain extends Message<Domain> {
  /**
   * Globally unique domain name.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * Display name.
   *
   * @generated from field: string name = 2;
   */
  name = "";

  constructor(data?: PartialMessage<Domain>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.Domain";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Domain {
    return new Domain().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Domain {
    return new Domain().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Domain {
    return new Domain().fromJsonString(jsonString, options);
  }

  static equals(a: Domain | PlainMessage<Domain> | undefined, b: Domain | PlainMessage<Domain> | undefined): boolean {
    return proto3.util.equals(Domain, a, b);
  }
}

/**
 * Represents a list of domains.
 *
 * @generated from message flyteidl.admin.GetDomainsResponse
 */
export class GetDomainsResponse extends Message<GetDomainsResponse> {
  /**
   * @generated from field: repeated flyteidl.admin.Domain domains = 1;
   */
  domains: Domain[] = [];

  constructor(data?: PartialMessage<GetDomainsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetDomainsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "domains", kind: "message", T: Domain, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDomainsResponse {
    return new GetDomainsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDomainsResponse {
    return new GetDomainsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDomainsResponse {
    return new GetDomainsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetDomainsResponse | PlainMessage<GetDomainsResponse> | undefined, b: GetDomainsResponse | PlainMessage<GetDomainsResponse> | undefined): boolean {
    return proto3.util.equals(GetDomainsResponse, a, b);
  }
}

/**
 * Top-level namespace used to classify different entities like workflows and executions.
 *
 * @generated from message flyteidl.admin.Project
 */
export class Project extends Message<Project> {
  /**
   * Globally unique project name.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * Display name.
   *
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: repeated flyteidl.admin.Domain domains = 3;
   */
  domains: Domain[] = [];

  /**
   * @generated from field: string description = 4;
   */
  description = "";

  /**
   * Leverage Labels from flyteidl.admin.common.proto to
   * tag projects with ownership information.
   *
   * @generated from field: flyteidl.admin.Labels labels = 5;
   */
  labels?: Labels;

  /**
   * @generated from field: flyteidl.admin.Project.ProjectState state = 6;
   */
  state = Project_ProjectState.ACTIVE;

  /**
   * Optional, org key applied to the resource.
   *
   * @generated from field: string org = 7;
   */
  org = "";

  constructor(data?: PartialMessage<Project>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.Project";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "domains", kind: "message", T: Domain, repeated: true },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "labels", kind: "message", T: Labels },
    { no: 6, name: "state", kind: "enum", T: proto3.getEnumType(Project_ProjectState) },
    { no: 7, name: "org", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Project {
    return new Project().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Project {
    return new Project().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Project {
    return new Project().fromJsonString(jsonString, options);
  }

  static equals(a: Project | PlainMessage<Project> | undefined, b: Project | PlainMessage<Project> | undefined): boolean {
    return proto3.util.equals(Project, a, b);
  }
}

/**
 * The state of the project is used to control its visibility in the UI and validity.
 *
 * @generated from enum flyteidl.admin.Project.ProjectState
 */
export enum Project_ProjectState {
  /**
   * By default, all projects are considered active.
   *
   * @generated from enum value: ACTIVE = 0;
   */
  ACTIVE = 0,

  /**
   * Archived projects are no longer visible in the UI and no longer valid.
   *
   * @generated from enum value: ARCHIVED = 1;
   */
  ARCHIVED = 1,

  /**
   * System generated projects that aren't explicitly created or managed by a user.
   *
   * @generated from enum value: SYSTEM_GENERATED = 2;
   */
  SYSTEM_GENERATED = 2,

  /**
   * System archived projects that aren't explicitly archived by a user.
   *
   * @generated from enum value: SYSTEM_ARCHIVED = 3;
   */
  SYSTEM_ARCHIVED = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Project_ProjectState)
proto3.util.setEnumType(Project_ProjectState, "flyteidl.admin.Project.ProjectState", [
  { no: 0, name: "ACTIVE" },
  { no: 1, name: "ARCHIVED" },
  { no: 2, name: "SYSTEM_GENERATED" },
  { no: 3, name: "SYSTEM_ARCHIVED" },
]);

/**
 * Represents a list of projects.
 * See :ref:`ref_flyteidl.admin.Project` for more details
 *
 * @generated from message flyteidl.admin.Projects
 */
export class Projects extends Message<Projects> {
  /**
   * @generated from field: repeated flyteidl.admin.Project projects = 1;
   */
  projects: Project[] = [];

  /**
   * In the case of multiple pages of results, the server-provided token can be used to fetch the next page
   * in a query. If there are no more results, this value will be empty.
   *
   * @generated from field: string token = 2;
   */
  token = "";

  constructor(data?: PartialMessage<Projects>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.Projects";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "projects", kind: "message", T: Project, repeated: true },
    { no: 2, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Projects {
    return new Projects().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Projects {
    return new Projects().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Projects {
    return new Projects().fromJsonString(jsonString, options);
  }

  static equals(a: Projects | PlainMessage<Projects> | undefined, b: Projects | PlainMessage<Projects> | undefined): boolean {
    return proto3.util.equals(Projects, a, b);
  }
}

/**
 * Request to retrieve a list of projects matching specified filters. 
 * See :ref:`ref_flyteidl.admin.Project` for more details
 *
 * @generated from message flyteidl.admin.ProjectListRequest
 */
export class ProjectListRequest extends Message<ProjectListRequest> {
  /**
   * Indicates the number of projects to be returned.
   * +required
   *
   * @generated from field: uint32 limit = 1;
   */
  limit = 0;

  /**
   * In the case of multiple pages of results, this server-provided token can be used to fetch the next page
   * in a query.
   * +optional
   *
   * @generated from field: string token = 2;
   */
  token = "";

  /**
   * Indicates a list of filters passed as string.
   * More info on constructing filters : <Link>
   * +optional
   *
   * @generated from field: string filters = 3;
   */
  filters = "";

  /**
   * Sort ordering.
   * +optional
   *
   * @generated from field: flyteidl.admin.Sort sort_by = 4;
   */
  sortBy?: Sort;

  /**
   * Optional, org filter applied to list project requests.
   *
   * @generated from field: string org = 5;
   */
  org = "";

  constructor(data?: PartialMessage<ProjectListRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.ProjectListRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "limit", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 2, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "filters", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "sort_by", kind: "message", T: Sort },
    { no: 5, name: "org", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProjectListRequest {
    return new ProjectListRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProjectListRequest {
    return new ProjectListRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProjectListRequest {
    return new ProjectListRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ProjectListRequest | PlainMessage<ProjectListRequest> | undefined, b: ProjectListRequest | PlainMessage<ProjectListRequest> | undefined): boolean {
    return proto3.util.equals(ProjectListRequest, a, b);
  }
}

/**
 * Adds a new user-project within the Flyte deployment.
 * See :ref:`ref_flyteidl.admin.Project` for more details
 *
 * @generated from message flyteidl.admin.ProjectRegisterRequest
 */
export class ProjectRegisterRequest extends Message<ProjectRegisterRequest> {
  /**
   * +required
   *
   * @generated from field: flyteidl.admin.Project project = 1;
   */
  project?: Project;

  constructor(data?: PartialMessage<ProjectRegisterRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.ProjectRegisterRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project", kind: "message", T: Project },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProjectRegisterRequest {
    return new ProjectRegisterRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProjectRegisterRequest {
    return new ProjectRegisterRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProjectRegisterRequest {
    return new ProjectRegisterRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ProjectRegisterRequest | PlainMessage<ProjectRegisterRequest> | undefined, b: ProjectRegisterRequest | PlainMessage<ProjectRegisterRequest> | undefined): boolean {
    return proto3.util.equals(ProjectRegisterRequest, a, b);
  }
}

/**
 * Purposefully empty, may be updated in the future.
 *
 * @generated from message flyteidl.admin.ProjectRegisterResponse
 */
export class ProjectRegisterResponse extends Message<ProjectRegisterResponse> {
  constructor(data?: PartialMessage<ProjectRegisterResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.ProjectRegisterResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProjectRegisterResponse {
    return new ProjectRegisterResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProjectRegisterResponse {
    return new ProjectRegisterResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProjectRegisterResponse {
    return new ProjectRegisterResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ProjectRegisterResponse | PlainMessage<ProjectRegisterResponse> | undefined, b: ProjectRegisterResponse | PlainMessage<ProjectRegisterResponse> | undefined): boolean {
    return proto3.util.equals(ProjectRegisterResponse, a, b);
  }
}

/**
 * Purposefully empty, may be updated in the future.
 *
 * @generated from message flyteidl.admin.ProjectUpdateResponse
 */
export class ProjectUpdateResponse extends Message<ProjectUpdateResponse> {
  constructor(data?: PartialMessage<ProjectUpdateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.ProjectUpdateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProjectUpdateResponse {
    return new ProjectUpdateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProjectUpdateResponse {
    return new ProjectUpdateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProjectUpdateResponse {
    return new ProjectUpdateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ProjectUpdateResponse | PlainMessage<ProjectUpdateResponse> | undefined, b: ProjectUpdateResponse | PlainMessage<ProjectUpdateResponse> | undefined): boolean {
    return proto3.util.equals(ProjectUpdateResponse, a, b);
  }
}

/**
 * @generated from message flyteidl.admin.ProjectGetRequest
 */
export class ProjectGetRequest extends Message<ProjectGetRequest> {
  /**
   * Indicates a unique project.
   * +required
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * Optional, org key applied to the resource.
   *
   * @generated from field: string org = 2;
   */
  org = "";

  constructor(data?: PartialMessage<ProjectGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.ProjectGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "org", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProjectGetRequest {
    return new ProjectGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProjectGetRequest {
    return new ProjectGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProjectGetRequest {
    return new ProjectGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ProjectGetRequest | PlainMessage<ProjectGetRequest> | undefined, b: ProjectGetRequest | PlainMessage<ProjectGetRequest> | undefined): boolean {
    return proto3.util.equals(ProjectGetRequest, a, b);
  }
}

/**
 * Error returned for inactive projects
 *
 * @generated from message flyteidl.admin.InactiveProject
 */
export class InactiveProject extends Message<InactiveProject> {
  /**
   * Indicates a unique project.
   * +required
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * Optional, org key applied to the resource.
   *
   * @generated from field: string org = 2;
   */
  org = "";

  constructor(data?: PartialMessage<InactiveProject>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.InactiveProject";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "org", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InactiveProject {
    return new InactiveProject().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InactiveProject {
    return new InactiveProject().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InactiveProject {
    return new InactiveProject().fromJsonString(jsonString, options);
  }

  static equals(a: InactiveProject | PlainMessage<InactiveProject> | undefined, b: InactiveProject | PlainMessage<InactiveProject> | undefined): boolean {
    return proto3.util.equals(InactiveProject, a, b);
  }
}

