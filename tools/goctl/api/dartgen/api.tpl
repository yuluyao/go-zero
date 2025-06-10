// dart format width=80
// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

import 'dart:io';

import 'package:dio/dio.dart';

import 'entity.dart';

class {{ highCamelCase .Service.Name }} {

  static {{ highCamelCase .Service.Name }} create({
    required String baseUrl,
    List<Interceptor>? interceptors,
  }) {
    final options = BaseOptions(
      baseUrl: baseUrl,
      connectTimeout: const Duration(seconds: 30),
      sendTimeout: const Duration(seconds: 30),
      receiveTimeout: const Duration(seconds: 30),
    );
    final dio = Dio(options);
    if (interceptors != null) {
      dio.interceptors.addAll(interceptors);
    }
    return {{ highCamelCase .Service.Name }}._(dio: dio);
  }

  final Dio dio;

  {{ highCamelCase .Service.Name }}._({required this.dio});

  /// 返回非 null 时，代表操作成功
  ///
  /// 1. 请求出错
  ///     return null
  /// 2. 请求成功 -> 业务出错
  ///     return null
  /// 3. 请求成功 -> 业务正常
  ///     - 有数据，return `Map<String, dynamic>` or `List<dynamic>`
  ///     - 无数据，return `String`，代表操作成功
  dynamic unwrapResponse(Response resp) {
    // 请求出错
    if (resp.statusCode != HttpStatus.ok) {
      return null;
    }
    Map<String, dynamic> respData = resp.data;
    final {'code': code, 'msg': msg, 'data': data} = respData;
    // 请求成功 -> 业务出错
    if (code != 0) {
      return null;
    }
    // 请求成功 -> 业务正常
    if (data == null) {
      return msg;
    } else {
      return data;
    }
  }

{{ range .Service.Groups }}
  {{ range .Routes }}
    Future<{{if .ResponseType}}{{.ResponseType.Name}}?{{else}}String?{{end}}> {{ lowCamelCase .Handler }}() async {
      final resp = await dio.{{lowCamelCase .Method}}('{{.Path}}');
      final data = unwrapResponse(resp);
      if (data == null) {
        return null;
      }
      return {{if .ResponseType}}{{.ResponseType.Name}}.fromJson(data){{else}}data{{end}};
    }
  {{end}}
{{end}}

}