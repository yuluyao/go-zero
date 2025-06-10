// dart format width=80
// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

import 'package:freezed_annotation/freezed_annotation.dart';

part 'entity.freezed.dart';
part 'entity.g.dart';

{{ range .APISpec.Types}}
  @freezed
  sealed class {{.Name}} with _${{.Name}} {
    const factory {{.Name}}(
      {{if .Members}}{
        {{range .Members}}
          {{if .Comment}}/// {{.Comment}}
          {{end}}
          {{if isNullableType .Type.Name}}
            @JsonKey(name: '{{getPropertyFromMember .}}') {{ specTypeToDart .Type }} {{lowCamelCase .Name}},
          {{else}}
            @JsonKey(name: '{{getPropertyFromMember .}}') required {{ specTypeToDart .Type }} {{lowCamelCase .Name}},
          {{end}}
        {{end}}
      }
      {{end}}) = _{{.Name}};

    factory {{.Name}}.fromJson(Map<String, dynamic> json) => _${{.Name}}FromJson(json);

    {{ range $.InnerClassList}}
    {{.}}
    {{end}}
  }
{{end}}