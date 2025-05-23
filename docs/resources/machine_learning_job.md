---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "grafana_machine_learning_job Resource - terraform-provider-grafana"
subcategory: "Machine Learning"
description: |-
  A job defines the queries and model parameters for a machine learning task.
  See the Grafana Cloud docs https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/dynamic-alerting/forecasting/config/ for more information
  on available hyperparameters for use in the hyper_params field.
---

# grafana_machine_learning_job (Resource)

A job defines the queries and model parameters for a machine learning task.

See [the Grafana Cloud docs](https://grafana.com/docs/grafana-cloud/alerting-and-irm/machine-learning/dynamic-alerting/forecasting/config/) for more information
on available hyperparameters for use in the `hyper_params` field.

## Example Usage

### Basic Forecast

This forecast uses a Prometheus datasource, where the source query is defined in the `expr` field of the `query_params` attribute.

Other datasources are supported, but the structure `query_params` may differ.

```terraform
resource "grafana_data_source" "foo" {
  type                = "prometheus"
  name                = "prometheus-ds-test"
  uid                 = "prometheus-ds-test-uid"
  url                 = "https://my-instance.com"
  basic_auth_enabled  = true
  basic_auth_username = "username"

  json_data_encoded = jsonencode({
    httpMethod        = "POST"
    prometheusType    = "Mimir"
    prometheusVersion = "2.4.0"
  })

  secure_json_data_encoded = jsonencode({
    basicAuthPassword = "password"
  })
}

resource "grafana_machine_learning_job" "test_job" {
  name            = "Test Job"
  metric          = "tf_test_job"
  datasource_type = "prometheus"
  datasource_uid  = grafana_data_source.foo.uid
  query_params = {
    expr = "grafanacloud_grafana_instance_active_user_count"
  }
}
```

### Tuned Forecast

This forecast has tuned hyperparameters to improve the accuracy of the model.

```terraform
resource "grafana_data_source" "foo" {
  type                = "prometheus"
  name                = "prometheus-ds-test"
  uid                 = "prometheus-ds-test-uid"
  url                 = "https://my-instance.com"
  basic_auth_enabled  = true
  basic_auth_username = "username"

  json_data_encoded = jsonencode({
    httpMethod        = "POST"
    prometheusType    = "Mimir"
    prometheusVersion = "2.4.0"
  })

  secure_json_data_encoded = jsonencode({
    basicAuthPassword = "password"
  })
}

resource "grafana_machine_learning_job" "test_job" {
  name            = "Test Job"
  metric          = "tf_test_job"
  datasource_type = "prometheus"
  datasource_uid  = grafana_data_source.foo.uid
  query_params = {
    expr = "grafanacloud_grafana_instance_active_user_count"
  }
  hyper_params = {
    daily_seasonality  = 15
    weekly_seasonality = 10
  }
  custom_labels = {
    example_label = "example_value"
  }
}
```

### Rescaled Forecast

This forecast has had the data transformed using a power transformation in order to avoid negative lower predictions.

```terraform
resource "grafana_data_source" "foo" {
  type                = "prometheus"
  name                = "prometheus-ds-test"
  uid                 = "prometheus-ds-test-uid"
  url                 = "https://my-instance.com"
  basic_auth_enabled  = true
  basic_auth_username = "username"

  json_data_encoded = jsonencode({
    httpMethod        = "POST"
    prometheusType    = "Mimir"
    prometheusVersion = "2.4.0"
  })

  secure_json_data_encoded = jsonencode({
    basicAuthPassword = "password"
  })
}

resource "grafana_machine_learning_job" "test_job" {
  name            = "Test Job"
  metric          = "tf_test_job"
  datasource_type = "prometheus"
  datasource_uid  = grafana_data_source.foo.uid
  query_params = {
    expr = "grafanacloud_grafana_instance_active_user_count"
  }
  hyper_params = {
    transformation_id = "power"
  }
}
```

### Forecast with Holidays

This forecast has holidays which will be taken into account when training the model.

```terraform
resource "grafana_data_source" "foo" {
  type                = "prometheus"
  name                = "prometheus-ds-test"
  uid                 = "prometheus-ds-test-uid"
  url                 = "https://my-instance.com"
  basic_auth_enabled  = true
  basic_auth_username = "username"

  json_data_encoded = jsonencode({
    httpMethod        = "POST"
    prometheusType    = "Mimir"
    prometheusVersion = "2.4.0"
  })

  secure_json_data_encoded = jsonencode({
    basicAuthPassword = "password"
  })
}

resource "grafana_machine_learning_holiday" "test_holiday" {
  name = "Test Holiday"
  custom_periods {
    name       = "First of January"
    start_time = "2023-01-01T00:00:00Z"
    end_time   = "2023-01-02T00:00:00Z"
  }
}

resource "grafana_machine_learning_job" "test_job" {
  name            = "Test Job"
  metric          = "tf_test_job"
  datasource_type = "prometheus"
  datasource_uid  = grafana_data_source.foo.uid
  query_params = {
    expr = "grafanacloud_grafana_instance_active_user_count"
  }
  holidays = [
    grafana_machine_learning_holiday.test_holiday.id
  ]
}
```

## Import

Import is supported using the following syntax:

```shell
terraform import grafana_machine_learning_job.name "{{ id }}"
```
