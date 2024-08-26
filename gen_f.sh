dp-cli gen --connstr "postgresql://things:things2024@localhost:5432/thingsdb?sslmode=disable" \
--tables=f_alarm,f_history_alarm,f_event --model_naming "{{ toUpperCamelCase ( replace . \"f_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"f_\" \"\") }}" \
--module alarm-service

