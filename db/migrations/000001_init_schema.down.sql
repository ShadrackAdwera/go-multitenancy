-- Drop foreign keys first
ALTER TABLE "user_group" DROP CONSTRAINT "user_group_group_id_fkey";
ALTER TABLE "user_group" DROP CONSTRAINT "user_group_user_id_fkey";
ALTER TABLE "group_policy" DROP CONSTRAINT "group_policy_policy_id_fkey";
ALTER TABLE "group_policy" DROP CONSTRAINT "group_policy_group_id_fkey";
ALTER TABLE "profile" DROP CONSTRAINT "profile_group_id_fkey";
ALTER TABLE "profile" DROP CONSTRAINT "profile_user_id_fkey";
ALTER TABLE "permission" DROP CONSTRAINT "permission_policy_id_fkey";
ALTER TABLE "policy" DROP CONSTRAINT "policy_group_id_fkey";
ALTER TABLE "user" DROP CONSTRAINT "user_tenant_id_fkey";

-- Drop indexes
DROP INDEX "user_group_group_id_user_id_idx";
DROP INDEX "group_policy_group_id_policy_id_idx";
DROP INDEX "profile_id_user_id_idx";
DROP INDEX "permission_id_policy_id_idx";
DROP INDEX "user_id_tenant_id_idx";

-- Drop tables in reverse order
DROP TABLE "user_group";
DROP TABLE "group_policy";
DROP TABLE "profile";
DROP TABLE "permission";
DROP TABLE "policy";
DROP TABLE "group";
DROP TABLE "user";
DROP TABLE "tenant";
