
id, err := tpgresource.ReplaceVars(d, config, "{{env_id}}/keyvaluemaps/{{name}}")
if err != nil {
	return fmt.Errorf("Error constructing id: %s", err)
}
d.SetId(id)

log.Printf("[DEBUG] Finished creating EnvironmentKeyvaluemaps %q: %#v", d.Id(), res)

return resourceApigeeEnvironmentKeyvaluemapsRead(d, meta)