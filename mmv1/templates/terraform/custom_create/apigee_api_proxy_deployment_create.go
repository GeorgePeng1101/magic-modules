config := meta.(*transport_tpg.Config)
userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
if err != nil {
	return err
}

url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/environments/{{environment}}/apis/{{proxy_name}}/revisions/{{revision}}/deployments?override=true&serviceAccount={{service_account}}")
if err != nil {
	return err
}

log.Printf("[DEBUG] Creating new ApiProxyDeployment at %s", url)
billingProject := ""

// err == nil indicates that the billing_project value was found
if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
	billingProject = bp
}

res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
	Config:    config,
	Method:    "POST",
	Project:   billingProject,
	RawURL:    url,
	UserAgent: userAgent,
	Timeout:   d.Timeout(schema.TimeoutCreate),
})
if err != nil {
	return fmt.Errorf("Error creating ApiProxyDeployment: %s", err)
}

// Store the ID now
id, err := tpgresource.ReplaceVars(d, config, "organizations/{{org_id}}/environments/{{environment}}/apis/{{proxy_name}}/revisions/{{revision}}/deployments")
if err != nil {
	return fmt.Errorf("Error constructing id: %s", err)
}
d.SetId(id)

log.Printf("[DEBUG] Finished creating ApiProxyDeployment %q: %#v", d.Id(), res)

return resourceApigeeApiProxyDeploymentRead(d, meta)